package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"go.uber.org/zap"
)

var (
	commands = []Command{
		{
			Name:    "/start",
			Handler: startCommand,
		},
	}
)

type TgBot struct {
	tgBot   *bot.Bot
	service *services.Services
}

type Command struct {
	Name    string
	Handler func(ctx context.Context, b *bot.Bot, update *models.Update)
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

func (b TgBot) Start(ctx context.Context) {
	b.Register()
	b.tgBot.Start(ctx)
}

func (b TgBot) Register() {
	for _, c := range commands {
		b.tgBot.RegisterHandler(bot.HandlerTypeMessageText, c.Name, bot.MatchTypeExact, c.Handler)
	}
	zap.S().Debug("Commands registered")
}
