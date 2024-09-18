package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ndfz/solana-nft-notify-bot/internal/worker"
)

// TODO: make a clear message
func (tg TgBot) notify(ctx context.Context) {
	for {
		activity := <-worker.ActivityUpdates
		users, _ := tg.service.User.GetByCollectionSymbol(activity.CollectionSymbol)

		for _, user := range users {
			message := fmt.Sprintf(
				"📢 *New NFT Sale Alert!*\n\n"+
					"*Collection:* %s\n"+
					"*NFT Token:* %s\n"+
					"*Seller:* %s\n"+
					"*Buyer:* %s\n"+
					"*Price:* %.2f SOL\n\n"+
					"*Transaction Signature:* %s\n"+
					"*Collection Symbol:* %s\n\n"+
					"🔗 [View on Magic Eden](https://magiceden.io/marketplace/%s)",
				activity.Collection,
				activity.TokenMint,
				activity.Seller,
				activity.Buyer,
				activity.Price,
				activity.Signature,
				activity.CollectionSymbol,
				activity.CollectionSymbol,
			)

			if activity.Image != "" {
				_, err := tg.tgBot.SendPhoto(ctx, &bot.SendPhotoParams{
					ChatID: user.TelegramID,
					// FIX: Bad Request: wrong file identifier/HTTP URL specified
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
