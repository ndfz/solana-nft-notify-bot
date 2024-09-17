package storage

type CollectionDTO struct {
	TelegramID int64  `json:"telegram_id"`
	Symbol     string `json:"symbol"`
}

type Collection struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
}
