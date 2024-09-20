package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

func helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	zap.S().Debugf("/help command called from: %d (%s)", update.Message.From.ID, update.Message.From.Username)

	helpMessage := `🤖 *Bot Command Reference*

Here are the commands you can use:

/start - *Initialize the bot*
/help - *Display this help message*
/addcollection - *Add a new collection*
/removecollection - *Remove an existing collection*
/listcollections - *List all your collections*

For more information, visit the [GitHub](https://github.com/ndfz/solana-nft-notify-bot).
`

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      helpMessage,
		ParseMode: models.ParseModeMarkdownV1,
	})
}
