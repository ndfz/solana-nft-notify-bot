package telegram

import (
	"github.com/ndfz/solana-nft-notify-bot/internal/worker"
	"go.uber.org/zap"
)

func notify() {
	for {
		activity := <-worker.ActivityUpdates
		zap.S().Info("notify: ", activity)
	}
}
