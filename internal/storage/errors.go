package storage

import "errors"

var (
	ErrUserExists         = errors.New("user already exists")
	ErrNoCollectionsFound = errors.New("no collections found for the user")
)
