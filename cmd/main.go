package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-telegram/bot"
	_ "github.com/mattn/go-sqlite3"
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

	logger := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
		logger.Debug("development mode")
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	zap.S().Info("starting solana-nft-notify-bot")

	storage, err := storage.New(config.DatabaseName)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	defer storage.DB.Close()
	zap.S().Info("connected to SQLite database")
	err = storage.CreateTables()
	if err != nil {
		panic("failed to create tables: " + err.Error())
	}
	zap.S().Info("created tables")

	// opts := bot.Option{
	//   bot.WithDefaultHandler()
	// }

	magiceden := magiceden.New(config.MagicEdenEndpoint)

	userRepository := user.New(storage.DB)
	collectionRepository := collection.New(storage.DB)

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

	tgBot, err := bot.New(config.TgBotToken)
	if err != nil {
		panic("failed to create bot: " + err.Error())
	}

	zap.S().Info("started bot")
	telegram.New(tgBot, services).Start(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
