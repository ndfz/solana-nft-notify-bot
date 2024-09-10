package services

import (
	"github.com/ndfz/solana-nft-notify-bot/internal/config"
)

type Services struct {
	Config *config.Config
}

func New(
	cfg *config.Config) *Services {
	return &Services{
		Config: cfg,
	}
}
