package coingecko

import (
	"context"
	"encoding/json"
	"gitlab.com/hotelian-company/challenge/config"
	"gitlab.com/hotelian-company/challenge/pkg/http"
	"time"
)

const (
	baseURL = "https://api.coingecko.com/api/v3/exchange_rates"
)

type CoinGecko struct {
	timeout time.Duration
}

type coin struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

type rateResponse struct {
	Rates map[string]coin `json:"rates"`
}

func NewCoinGecko() *CoinGecko {
	return &CoinGecko{}
}

func (c *CoinGecko) GetRate(ctx context.Context, from, to string) (float64, error) {
	url := makeURL()

	res, err := http.NewRequest(http.MethodGET, url, nil)
	if err != nil {
		return 0, err
	}

	response := &rateResponse{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return 0, err
	}

	result, err := compareRates(response.Rates, from, to)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *CoinGecko) GetName() string {
	return config.COINGECKO
}

func compareRates(rates map[string]coin, from, to string) (result float64, err error) {
	var fromVal, toVal float64
	for symbol, r := range rates {
		if symbol == from {
			fromVal = r.Value
		}

		if symbol == to {
			toVal = r.Value
		}
	}

	if fromVal != 0 && toVal != 0 {
		result = fromVal / toVal
	}

	return result, err
}

func makeURL() string {
	return baseURL
}
