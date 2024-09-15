package worker

import (
	"time"

	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"go.uber.org/zap"
)

var (
	targetAction = "buyNow"
)

type Worker struct {
	services *services.Services
}

func New(services *services.Services) Worker {
	return Worker{
		services: services,
	}
}

func (w Worker) Run() {
	zap.S().Info("starting worker")
	// TODO: get collections from database
	//
	// this is just for testing
	collections := []string{"y00ts", "retardio_cousins"}

	for {
		for _, c := range collections {
			result := w.services.Magiceden.GetActivitiesOfCollection(c)
			for _, r := range result {
				if r.Type == targetAction {
					zap.S().Info(r)
				}
			}
			zap.S().Debug("sleeping " + w.services.Config.CollectionSleep.String())
			time.Sleep(w.services.Config.CollectionSleep)
		}
		zap.S().Debug("sleeping " + w.services.Config.CycleSleep.String())
		time.Sleep(w.services.Config.CycleSleep)
	}
}
