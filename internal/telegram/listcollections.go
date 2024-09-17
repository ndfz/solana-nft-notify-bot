package telegram

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"go.uber.org/zap"
)

func listCollectionsHandler(ctx context.Context, b *bot.Bot, update *models.Update, service *services.Services) {
	collections, err := service.Collection.GetByTelegramID(update.Message.From.ID)
	if err != nil {
		zap.S().Errorf("Error getting collections: %v", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Error getting collections!",
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Your collections: %v", collections),
	})
}
