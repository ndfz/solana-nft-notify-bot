package impl

import (
	"database/sql"

	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func New(db *sql.DB) UserRepositoryImpl {
	return UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) SaveUser(user storage.UserDTO) error {
	return nil
}
