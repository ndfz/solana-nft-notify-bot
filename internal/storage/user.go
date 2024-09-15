package storage

type UserDTO struct {
	ID         string `json:"id"`
	TelegramID string `json:"telegram_id"`
}

type User struct {
	ID         string `json:"id"`
	TelegramID string `json:"telegram_id"`
	CreatedAt  string `json:"created_at"`
}
