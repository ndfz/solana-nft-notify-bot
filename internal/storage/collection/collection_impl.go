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
	var collectionID uuid.UUID

	err := r.db.QueryRow("SELECT id FROM collections WHERE symbol = ?", collection.Symbol).Scan(&collectionID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing collection: %v", err)
	}

	if err == sql.ErrNoRows {
		collectionID = uuid.New()
		_, err = r.db.Exec("INSERT INTO collections (id, symbol) VALUES (?, ?)", collectionID, collection.Symbol)
		if err != nil {
			return fmt.Errorf("failed to save collection: %v", err)
		}
		zap.S().Infof("collection saved: %s", collection.Symbol)
	} else {
		zap.S().Infof("collection already exists: %s", collection.Symbol)
	}

	var userID uuid.UUID
	err = r.db.QueryRow("SELECT id FROM users WHERE telegram_id = ?", collection.TelegramID).Scan(&userID)
	if err != nil {
		return fmt.Errorf("failed to find user by telegram_id: %v", err)
	}

	var userCollectionID uuid.UUID
	err = r.db.QueryRow("SELECT id FROM users_collections WHERE user_id = ? AND collections_id = ?", userID, collectionID).Scan(&userCollectionID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing user-collection relationship: %v", err)
	}

	if err == sql.ErrNoRows {
		uuidUserCollection := uuid.New()
		_, err = r.db.Exec("INSERT INTO users_collections (id, user_id, collections_id) VALUES (?, ?, ?)", uuidUserCollection, userID, collectionID)
		if err != nil {
			return fmt.Errorf("failed to save user-collection relationship: %v", err)
		}
		zap.S().Infof("user-Collection relationship saved: User %d - Collection %s", collection.TelegramID, collection.Symbol)
	} else {
		zap.S().Infof("user-Collection relationship already exists: User %d - Collection %s", collection.TelegramID, collection.Symbol)
	}

	return nil
}

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

func (r CollectionRepositoryImpl) GetByTelegramID(telegramID int64) ([]storage.Collection, error) {
	query := `
		SELECT c.id, c.symbol
		FROM collections c
		INNER JOIN users_collections uc ON c.id = uc.collections_id
		INNER JOIN users u ON uc.user_id = u.id
		WHERE u.telegram_id = $1;
	`

	rows, err := r.db.Query(query, telegramID)
	if err != nil {
		return nil, fmt.Errorf("failed to get collections: %v", err)
	}
	defer rows.Close()

	var collections []storage.Collection

	for rows.Next() {
		var collection storage.Collection
		if err := rows.Scan(&collection.ID, &collection.Symbol); err != nil {
			return nil, fmt.Errorf("failed to scan collection: %v", err)
		}
		collections = append(collections, collection)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %v", err)
	}

	if len(collections) == 0 {
		return nil, fmt.Errorf("no collections found for user %d", telegramID)
	}

	return collections, nil
}

func (r CollectionRepositoryImpl) DeleteBySymbol(symbol string) error {
	result, err := r.db.Exec("DELETE FROM collections WHERE symbol = ?", symbol)
	if err != nil {
		return fmt.Errorf("failed to delete collection with symbol %s: %v", symbol, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected for symbol %s: %v", symbol, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no collection found with symbol %s", symbol)
	}

	return nil
}
