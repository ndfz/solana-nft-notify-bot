package storage

type CollectionDTO struct {
	Symbol string `json:"symbol"`
}

type Collection struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
}
