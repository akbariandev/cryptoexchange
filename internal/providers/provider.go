package providers

import (
	"context"
	"errors"
	"github.com/akbarian.dev/cryptoexchange/config"
	"github.com/akbarian.dev/cryptoexchange/internal/providers/coingecko"
	"github.com/akbarian.dev/cryptoexchange/internal/providers/kucoin"
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
