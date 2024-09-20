package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/worker"
)

func (tg TgBot) notify(ctx context.Context) {
	for {
		activity := <-worker.ActivityUpdates
		users, _ := tg.service.User.GetByCollectionSymbol(activity.CollectionSymbol)

		for _, user := range users {
			message := fmt.Sprintf(
				"📢 *New NFT Sale Alert!*\n\n"+
					"🏷 *Collection:* %s\n"+
					"🖼 *NFT Token:* %s\n"+
					"👤 *Seller:* %s\n"+
					"🎉 *Buyer:* %s\n"+
					"💰 *Price:* %.3f SOL\n\n"+
					"🔗 *Transaction:* [%s](https://solscan.io/tx/%s)\n"+
					"🌐 *View Collection:* [%s](https://magiceden.io/marketplace/%s)\n",
				activity.Collection,
				activity.TokenMint,
				activity.Seller,
				activity.Buyer,
				activity.Price,
				activity.Signature,
				activity.Signature,
				activity.Collection,
				activity.CollectionSymbol,
			)

			if activity.Image != "" {
				_, err := tg.tgBot.SendPhoto(ctx, &bot.SendPhotoParams{
					ChatID:    user.TelegramID,
					Photo:     &models.InputFileString{Data: activity.Image},
					Caption:   message,
					ParseMode: models.ParseModeMarkdownV1,
				})
				if err != nil {
					log.Printf("Failed to send photo: %v", err)
				}
			} else {
				_, err := tg.tgBot.SendMessage(ctx, &bot.SendMessageParams{
					ChatID:    user.TelegramID,
					Text:      message,
					ParseMode: models.ParseModeMarkdownV1,
				})
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}
			}
		}
	}
}
