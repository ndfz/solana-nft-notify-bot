package services

import (
	"github.com/ndfz/solana-nft-notify-bot/internal/config"
	"github.com/ndfz/solana-nft-notify-bot/internal/magiceden"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
)

type Magiceden interface {
	GetActivitiesOfCollection(collectionName string) []magiceden.CollectionResponse
}

type Services struct {
	Config    *config.Config
	Storage   *storage.Storage
	Magiceden Magiceden
}

func New(
	cfg *config.Config,
	storage *storage.Storage,
	magiceden Magiceden,
) *Services {
	return &Services{
		Config:    cfg,
		Storage:   storage,
		Magiceden: magiceden,
	}
}
