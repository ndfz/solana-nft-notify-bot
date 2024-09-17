package collection

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ndfz/solana-nft-notify-bot/internal/storage"
	"go.uber.org/zap"
)

type CollectionRepositoryImpl struct {
	db *sql.DB
}

func New(db *sql.DB) CollectionRepositoryImpl {
	return CollectionRepositoryImpl{
		db: db,
	}
}

func (r CollectionRepositoryImpl) Save(collection storage.CollectionDTO) error {
	uuidCollection := uuid.New()

	_, err := r.db.Exec("INSERT INTO collections (id, symbol) VALUES (?, ?)", uuidCollection, collection.Symbol)
	if err != nil {
		return fmt.Errorf("failed to save collection: %v", err)
	}
	zap.S().Infof("Collection saved: %s", collection.Symbol)

	var userID uuid.UUID
	err = r.db.QueryRow("SELECT id FROM users WHERE telegram_id = ?", collection.TelegramID).Scan(&userID)
	if err != nil {
		return fmt.Errorf("failed to find user by telegram_id: %v", err)
	}

	uuidUserCollection := uuid.New()
	_, err = r.db.Exec("INSERT INTO users_collections (id, user_id, collections_id) VALUES (?, ?, ?)", uuidUserCollection, userID, uuidCollection)
	if err != nil {
		return fmt.Errorf("failed to save user-collection relationship: %v", err)
	}

	zap.S().Infof("User-Collection relationship saved: User %d - Collection %s", collection.TelegramID, collection.Symbol)

	return nil
}

// TODO: check this function
func (r CollectionRepositoryImpl) GetAll() ([]storage.Collection, error) {
	rows, err := r.db.Query("SELECT id, symbol FROM collections")
	if err != nil {
		return nil, fmt.Errorf("failed to get collections: %v", err)
	}
	defer rows.Close()

	var collections []storage.Collection
	for rows.Next() {
		var id string
		var symbol string
		if err := rows.Scan(&id, &symbol); err != nil {
			return nil, fmt.Errorf("failed to scan collection: %v", err)
		}
		collections = append(collections, storage.Collection{
			ID:     id,
			Symbol: symbol,
		})
	}
	return collections, nil
}

func (r CollectionRepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM collections WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete collection: %v", err)
	}
	zap.S().Infof("Collection deleted: %s", id)
	return nil
}
