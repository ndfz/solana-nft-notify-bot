package worker

import (
	"sort"
	"time"

	"github.com/ndfz/solana-nft-notify-bot/internal/magiceden"
	"github.com/ndfz/solana-nft-notify-bot/internal/services"
	"go.uber.org/zap"
)

var (
	ActivityUpdates = make(chan magiceden.CollectionResponse, 1)
	targetType      = "buyNow"
)

type Worker struct {
	services        *services.Services
	processedEvents map[string]time.Time
	eventTTL        time.Duration
}

func New(services *services.Services) Worker {
	return Worker{
		services:        services,
		processedEvents: make(map[string]time.Time),
		// TODO: make it configurable
		eventTTL: 60 * time.Minute,
	}
}

func (w Worker) Run() {
	for {
		collections, err := w.services.Collection.GetAll()
		if err != nil {
			zap.S().Error(err)
		}

		var allActivities []magiceden.CollectionResponse

		for _, c := range collections {
			zap.S().Debug("processing collection " + c.Symbol)
			result := w.services.Magiceden.GetActivitiesOfCollection(c.Symbol)
			for _, r := range result {
				if r.Type == targetType {
					if _, processed := w.processedEvents[r.Signature]; !processed {
						allActivities = append(allActivities, r)
						w.processedEvents[r.Signature] = time.Now()
					}
				}
			}
			zap.S().Debug("sleeping " + w.services.Config.CollectionSleep.String())
			time.Sleep(w.services.Config.CollectionSleep)
		}

		sort.Slice(allActivities, func(i, j int) bool {
			return allActivities[i].Signature > allActivities[j].Signature
		})

		for _, activity := range allActivities {
			zap.S().Debugf(
				"sending notification for %s, seller %s -> buyer %s, image %s",
				activity.Signature,
				activity.Seller,
				activity.Buyer,
				activity.Image,
			)
			ActivityUpdates <- activity
		}

		now := time.Now()
		for id, timestamp := range w.processedEvents {
			if now.Sub(timestamp) > w.eventTTL {
				delete(w.processedEvents, id)
			}
		}

		zap.S().Debug("sleeping " + w.services.Config.CycleSleep.String())
		time.Sleep(w.services.Config.CycleSleep)
	}
}
