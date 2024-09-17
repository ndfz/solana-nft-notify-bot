package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
	"go.uber.org/zap"
)

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update, service *services.Services) {
	dto := storage.UserDTO{
		TelegramID: update.Message.From.ID,
	}
	err := service.User.Save(dto)
	if err != nil {
		if err == storage.ErrUserExists {
			zap.S().Infof("user already exists: %d", update.Message.From.ID)
			return
		} else {
			zap.S().Errorf("error saving user: %v", err)
			return
		}
	}
}
