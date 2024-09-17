package storage

type UserDTO struct {
	TelegramID int64 `json:"telegram_id"`
}

type User struct {
	ID         string `json:"id"`
	TelegramID string `json:"telegram_id"`
	CreatedAt  string `json:"created_at"`
}
