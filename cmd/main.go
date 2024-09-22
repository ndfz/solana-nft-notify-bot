package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-telegram/bot"
	_ "github.com/lib/pq"
	"github.com/ndfz/solana-nft-notify-bot/internal/config"
	"github.com/ndfz/solana-nft-notify-bot/internal/magiceden"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage/collection"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage/user"
	"github.com/ndfz/solana-nft-notify-bot/internal/telegram"
	"github.com/ndfz/solana-nft-notify-bot/internal/worker"
	"go.uber.org/zap"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	var logger *zap.Logger

	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
		logger.Debug("development mode")
	} else {
		logger = zap.Must(zap.NewProduction())
		logger.Info("production mode")
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	zap.S().Info("starting solana-nft-notify-bot")

	storage, err := storage.New(config.DatabaseUrl)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	defer storage.Close()
	zap.S().Info("connected to postgresql")

	opts := []bot.Option{
		bot.WithMiddlewares(telegram.ShowCommandWithUserID),
	}

	magiceden := magiceden.New(config.MagicEdenEndpoint)

	userRepository := user.New(storage)
	collectionRepository := collection.New(storage)

	services := services.New(
		config,
		magiceden,
		userRepository,
		collectionRepository,
	)

	go worker.New(services).Run()
	zap.S().Info("started worker")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	tgBot, err := bot.New(config.TgBotToken, opts...)
	if err != nil {
		panic("failed to create bot: " + err.Error())
	}

	zap.S().Info("started bot")
	telegram.New(tgBot, services).Start(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
