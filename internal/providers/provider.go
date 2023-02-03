package providers

import (
	"context"
	"errors"
	"gitlab.com/hotelian-company/challenge/config"
	"gitlab.com/hotelian-company/challenge/internal/providers/coingecko"
	"gitlab.com/hotelian-company/challenge/internal/providers/kucoin"
)

type IProvider interface {
	GetRate(context.Context, string, string) (float64, error)
	GetName() string
}

func GetProvider(ctx context.Context, name string) (IProvider, error) {
	switch name {
	case config.KUCOIN:
		return kucoin.NewKuCoin(), nil
	case config.COINGECKO:
		return coingecko.NewCoinGecko(), nil
	}
	return nil, errors.New("provider not implemented")
}
