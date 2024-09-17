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
		zap.S().Errorf("Error saving user: %v", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Something went wrong!",
		})
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "You are registered!",
	})
}
