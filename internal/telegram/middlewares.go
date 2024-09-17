package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

func ShowCommandWithUserID(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			zap.S().Infof("%s  command called from: %d (%s)", update.Message.Text, update.Message.From.ID, update.Message.From.Username)
		}
		next(ctx, b, update)
	}
}
