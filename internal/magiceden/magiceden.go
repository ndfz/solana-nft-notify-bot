package magiceden

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type Magiceden struct {
	magicedenEndpoint string
}

func New(magicedenEndpoint string) *Magiceden {
	return &Magiceden{
		magicedenEndpoint: magicedenEndpoint,
	}
}

type CollectionResponse struct {
	Signature        string  `json:"signature"`
	Type             string  `json:"type"`
	TokenMint        string  `json:"tokenMint"`
	Collection       string  `json:"collection"`
	CollectionSymbol string  `json:"collectionSymbol"`
	Buyer            string  `json:"buyer"`
	Seller           string  `json:"seller"`
	Price            float64 `json:"price"`
	Image            string  `json:"image"`
}

func (m Magiceden) GetActivitiesOfCollection(collectionName string) []CollectionResponse {
	var data []CollectionResponse

	url := fmt.Sprintf("%s/collections/%s/activities?offset=0&limit=100", m.magicedenEndpoint, collectionName)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		zap.S().Errorf("error parsing JSON: %v", err)
		return nil
	}

	return data
}
