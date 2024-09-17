package services

import (
	"github.com/ndfz/solana-nft-notify-bot/internal/config"
	"github.com/ndfz/solana-nft-notify-bot/internal/magiceden"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
)

type Magiceden interface {
	GetActivitiesOfCollection(collectionName string) []magiceden.CollectionResponse
}

type (
	UserRepository interface {
		Save(user storage.UserDTO) error
	}
	CollectionRepository interface {
		Save(collection storage.CollectionDTO) error
		GetAll() ([]storage.Collection, error)
		GetByTelegramID(telegramID int64) ([]storage.Collection, error)
		DeleteBySymbol(symbol string) error
	}
)

type Services struct {
	Config     *config.Config
	Magiceden  Magiceden
	User       UserRepository
	Collection CollectionRepository
}

func New(
	cfg *config.Config,
	magiceden Magiceden,
	user UserRepository,
	collection CollectionRepository,
) *Services {
	return &Services{
		Config:     cfg,
		Magiceden:  magiceden,
		User:       user,
		Collection: collection,
	}
}
