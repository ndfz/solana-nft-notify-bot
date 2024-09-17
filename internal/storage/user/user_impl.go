package user

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
	"go.uber.org/zap"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func New(db *sql.DB) UserRepositoryImpl {
	return UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) Save(user storage.UserDTO) error {
	uuid := uuid.New()

	_, err := u.db.Exec("INSERT INTO users (id, telegram_id) VALUES (?, ?)", uuid, user.TelegramID)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}
	zap.S().Infof("User saved: %d", user.TelegramID)

	return nil
}
