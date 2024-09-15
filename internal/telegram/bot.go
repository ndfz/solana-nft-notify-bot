package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
)

type TgBot struct {
	tgBot   *bot.Bot
	service *services.Services
}

func New(
	bot *bot.Bot,
	service *services.Services,
) *TgBot {
	return &TgBot{
		tgBot:   bot,
		service: service,
	}
}

func (b TgBot) Run(ctx context.Context) {
	b.tgBot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Hello!",
		})
	})

	b.tgBot.Start(ctx)
}
