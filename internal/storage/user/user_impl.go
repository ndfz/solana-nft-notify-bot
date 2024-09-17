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
	var existingUserID uuid.UUID
	err := u.db.QueryRow("SELECT id FROM users WHERE telegram_id = ?", user.TelegramID).Scan(&existingUserID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking existing user: %v", err)
	}

	if err == nil {
		zap.S().Infof("user already exists: %d", user.TelegramID)
		return storage.ErrUserExists
	}

	uuid := uuid.New()
	_, err = u.db.Exec("INSERT INTO users (id, telegram_id) VALUES (?, ?)", uuid, user.TelegramID)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	zap.S().Infof("user saved: %d", user.TelegramID)
	return nil
}
