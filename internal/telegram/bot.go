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

func (tg TgBot) Start(ctx context.Context) {
	tg.Register()
	go notify()
	go tg.tgBot.Start(ctx)
}

func (tg TgBot) Register() {
	tg.tgBot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		startHandler(ctx, b, update, tg.service)
	})
	tg.tgBot.RegisterHandler(bot.HandlerTypeMessageText, "/addcollection", bot.MatchTypePrefix, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		addCollectionCommand(ctx, b, update, tg.service)
	})
	tg.tgBot.RegisterHandler(bot.HandlerTypeMessageText, "/removecollection", bot.MatchTypePrefix, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		removeCollectionCommand(ctx, b, update, tg.service)
	})
}
