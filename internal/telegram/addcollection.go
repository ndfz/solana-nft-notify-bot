package telegram

import (
	"context"
	"regexp"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
	"go.uber.org/zap"
)

func addCollectionCommand(ctx context.Context, b *bot.Bot, update *models.Update, service *services.Services) {
	dto := storage.CollectionDTO{
		TelegramID: update.Message.From.ID,
		Symbol:     getSymbolWithoutCommand(update.Message.Text),
	}
	err := service.Collection.Save(dto)
	if err != nil {
		zap.S().Errorf("Error saving collection: %v", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Error saving collection!",
		})
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Collection saved!",
	})
}

func getSymbolWithoutCommand(text string) string {
	re := regexp.MustCompile(`^/addcollection\s+(.+)$`)
	matches := re.FindStringSubmatch(text)
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}
