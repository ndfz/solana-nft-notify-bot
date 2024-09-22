package user

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
	"go.uber.org/zap"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UserRepositoryImpl {
	return UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) Save(user storage.UserDTO) error {
	var existingUserID uuid.UUID
	err := u.db.QueryRow("SELECT id FROM users WHERE telegram_id = $1", user.TelegramID).Scan(&existingUserID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking existing user: %v", err)
	}

	if err == nil {
		zap.S().Infof("user already exists: %d", user.TelegramID)
		return storage.ErrUserExists
	}

	_, err = u.db.Exec("INSERT INTO users (telegram_id) VALUES ($1)", user.TelegramID)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	zap.S().Infof("user saved: %d", user.TelegramID)
	return nil
}

func (u UserRepositoryImpl) GetByCollectionSymbol(symbol string) ([]storage.User, error) {
	query := `
		SELECT u.id, u.telegram_id, u.created_at
		FROM users u
		JOIN users_collections uc ON u.id = uc.user_id
		JOIN collections c ON uc.collections_id = c.id
		WHERE c.symbol = $1;
	`

	rows, err := u.db.Query(query, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to get users by collection symbol: %v", err)
	}
	defer rows.Close()

	var users []storage.User

	for rows.Next() {
		var user storage.User
		if err := rows.Scan(&user.ID, &user.TelegramID, &user.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate users: %v", err)
	}

	return users, nil
}
