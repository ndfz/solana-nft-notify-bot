package main

import (
	"os"

	"github.com/ndfz/solana-nft-notify-bot/internal/config"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"go.uber.org/zap"
)

func main() {
	config, err := config.New(nil)
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	logger := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
		logger.Debug("Development mode")
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	logger.Sugar().Info("Starting solana-nft-notify-bot")

	_ = services.New(config)
}
