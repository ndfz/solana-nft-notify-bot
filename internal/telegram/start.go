package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
)

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update, service *services.Services) {
	service.User.Save(storage.UserDTO{
		TelegramID: update.Message.From.ID,
	})
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "You are registered!",
	})
}
