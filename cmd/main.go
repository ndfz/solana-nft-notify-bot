package main

import (
	"github.com/ndfz/solana-nft-notify-bot/internal/config"
)

func main() {
	_, err := config.New(nil)
	if err != nil {
		panic(err)
	}
}
