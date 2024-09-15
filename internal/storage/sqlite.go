package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	DB *sql.DB
}

func New(databaseName string) (*Storage, error) {
	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	return &Storage{
		DB: db,
	}, nil
}

func (s Storage) CreateTables() error {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		telegram_id VARCHAR NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	collectionsTable := `
	CREATE TABLE IF NOT EXISTS collections (
		id UUID PRIMARY KEY,
		collection_name VARCHAR NOT NULL
	);`

	usersCollectionsTable := `
	CREATE TABLE IF NOT EXISTS users_collections (
		id UUID PRIMARY KEY,
		user_id UUID NOT NULL,
		collections_id UUID NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(collections_id) REFERENCES collections(id) ON DELETE CASCADE
	);`

	if _, err := s.DB.Exec(usersTable); err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	if _, err := s.DB.Exec(collectionsTable); err != nil {
		return fmt.Errorf("failed to create collections table: %v", err)
	}

	if _, err := s.DB.Exec(usersCollectionsTable); err != nil {
		return fmt.Errorf("failed to create users_collections table: %v", err)
	}

	return nil
}
