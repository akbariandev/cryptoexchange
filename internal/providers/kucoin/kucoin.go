package kucoin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/akbarian.dev/cryptoexchange/config"
	"github.com/akbarian.dev/cryptoexchange/pkg/http"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL = "https://api.kucoin.com/api/v1/market/stats"
)

type KuCoin struct {
	timeout time.Duration
}

type rateResponse struct {
	Code string `json:"code"`
	Data struct {
		Time             int64  `json:"time"`
		Symbol           string `json:"symbol"`
		Buy              string `json:"buy"`
		Sell             string `json:"sell"`
		ChangeRate       string `json:"changeRate"`
		ChangePrice      string `json:"changePrice"`
		High             string `json:"high"`
		Low              string `json:"low"`
		Vol              string `json:"vol"`
		VolValue         string `json:"volValue"`
		Last             string `json:"last"`
		AveragePrice     string `json:"averagePrice"`
		TakerFeeRate     string `json:"takerFeeRate"`
		MakerFeeRate     string `json:"makerFeeRate"`
		TakerCoefficient string `json:"takerCoefficient"`
		MakerCoefficient string `json:"makerCoefficient"`
	} `json:"data"`
}

type KuCoinAdapter struct {
	Provider *KuCoin
}

func NewKuCoin() *KuCoin {
	return &KuCoin{}
}

func (k *KuCoin) GetRate(ctx context.Context, from, to string) (float64, error) {
	url := makeURL(from, to)

	res, err := http.NewRequest(http.MethodGET, url, nil)
	if err != nil {
		return 0, err
	}

	response := &rateResponse{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(response.Data.ChangeRate, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

func (k *KuCoin) GetName() string {
	return config.KUCOIN
}

func makeURL(from, to string) string {
	return fmt.Sprintf("%s?symbol=%s-%s", baseURL, strings.ToUpper(strings.TrimSpace(from)), strings.ToUpper(strings.TrimSpace(to)))
}
