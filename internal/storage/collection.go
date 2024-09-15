package storage

type Collection struct {
	ID                string `json:"id"`
	CollectionName    string `json:"collection_name"`
	CollectionAddress string `json:"collection_address"`
	Marketplace       string `json:"marketplace"`
}
