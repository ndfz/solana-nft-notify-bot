package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

// TODO: implement this
func addCollectionCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	zap.S().Debugf("%s  command called from: %d (%s)", update.Message.Text, update.Message.From.ID, update.Message.From.Username)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Add collection!",
	})
}
