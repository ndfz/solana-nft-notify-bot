package telegram

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
	"go.uber.org/zap"
)

func listCollectionsHandler(ctx context.Context, b *bot.Bot, update *models.Update, service *services.Services) {
	collectionsFromDB, err := service.Collection.GetByTelegramID(update.Message.From.ID)
	if err != nil {
		if err == storage.ErrNoCollectionsFound {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "❌ You don't have any collections saved yet.",
			})
			return
		} else {
			zap.S().Errorf("Error getting collections: %v", err)
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "⚠️ An error occurred while retrieving your collections.",
			})
			return
		}
	}

	var collections []string
	for _, collection := range collectionsFromDB {
		collections = append(collections, collection.Symbol)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("📂 Here are your saved collections:\n%v", collections),
	})
}
