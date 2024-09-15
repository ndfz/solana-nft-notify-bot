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
	uuid := uuid.New()

	_, err := r.db.Exec("INSERT INTO collections (id, symbol) VALUES (?, ?)", uuid, collection.Symbol)
	if err != nil {
		return fmt.Errorf("failed to save collection: %v", err)
	}
	zap.S().Infof("Collection saved: %s", collection.Symbol)

	return nil
}
