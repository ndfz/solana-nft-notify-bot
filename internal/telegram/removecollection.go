package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"go.uber.org/zap"
)

func removeCollectionHandler(ctx context.Context, b *bot.Bot, update *models.Update, service *services.Services) {
	err := service.Collection.DeleteBySymbol(getTextWithoutCommand(update.Message.Text))
	if err != nil {
		zap.S().Errorf("Error removing collection: %v", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "⚠️ An error occurred while removing the collection.",
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "✅ Collection successfully removed! You will no longer receive updates for it.",
	})
}
