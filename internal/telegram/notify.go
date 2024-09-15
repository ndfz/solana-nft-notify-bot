// TODO: move to a new notify package
package telegram

import (
	"github.com/ndfz/solana-nft-notify-bot/internal/magiceden/worker"
	"go.uber.org/zap"
)

func notify() {
	for {
		waitChan := <-worker.NotifyChan
		zap.S().Info("notify: ", waitChan)
	}
}
