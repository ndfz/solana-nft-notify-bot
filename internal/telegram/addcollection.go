package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
	"go.uber.org/zap"
)

func addCollectionHandler(ctx context.Context, b *bot.Bot, update *models.Update, service *services.Services) {
	dto := storage.CollectionDTO{
		TelegramID: update.Message.From.ID,
		Symbol:     getTextWithoutCommand(update.Message.Text),
	}
	err := service.Collection.Save(dto)
	if err != nil {
		zap.S().Errorf("Error saving collection: %v", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Error saving collection!",
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Collection saved!",
	})
}
