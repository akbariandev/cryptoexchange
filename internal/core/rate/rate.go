package rate

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/hotelian-company/challenge/config"
	"gitlab.com/hotelian-company/challenge/internal/entity"
	"gitlab.com/hotelian-company/challenge/internal/providers"
	"gitlab.com/hotelian-company/challenge/pkg/logger"
	"go.uber.org/zap"
	"sort"
	"strings"
	"sync"
	"time"
)

func GetCurrenciesRate(ctx context.Context, currencies []string, to string) (rates []float64, err error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return rates, err
	}

	for _, from := range currencies {
		if ok := validateCurrency(cfg.Currencies, from); !ok {
			return rates, errors.New(fmt.Sprintf("currency %s not validated", from))
		}

		c := make(chan []float64)
		go func(from, to string, ch chan<- []float64) {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error(err.Error())
				}
			}()

			if err := getRateOfTwo(ctx, cfg.Providers, from, to, c); err != nil {
				panic(err)
			}
		}(from, to, c)

		rates = append(rates, mergeRates(<-c))
	}

	return rates, nil
}

func mergeRates(rates []float64) float64 {
	sort.Float64s(rates)
	for _, r := range rates {
		if r != 0 {
			return r
		}
	}

	return 0
}

func getRateOfTwo(ctx context.Context, cfgProviders map[string]entity.Provider, from, to string, ch chan<- []float64) error {

	var wg sync.WaitGroup
	results := make(chan float64)
	for providerName, providerConfig := range cfgProviders {
		if !providerConfig.Enable {
			logger.Logger.Warn("provider is not enable", zap.String("Provider", providerName))
			continue
		}

		d, _ := time.ParseDuration(providerConfig.Timeout)
		ctx2, _ := context.WithTimeout(ctx, d)
		provider, err := providers.GetProvider(ctx2, providerName)
		if err != nil {
			logger.Logger.Error(err.Error(), zap.String("Provider", providerName))
			continue
		}

		wg.Add(1)
		go func(from, to string) {
			defer wg.Done()
			result, err := provider.GetRate(ctx, from, to)
			if err != nil {
				logger.Logger.Error(err.Error(), zap.String("Provider", provider.GetName()))
				results <- 0
			}

			results <- result
		}(from, to)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var output []float64
	for result := range results {
		output = append(output, result)
	}

	ch <- output
	return nil
}

func validateCurrency(cfgCurrencies map[string]entity.Currency, currency string) bool {
	for name, config := range cfgCurrencies {
		if config.Enable {
			if strings.ToUpper(strings.TrimSpace(name)) == strings.ToUpper(strings.TrimSpace(currency)) {
				return true
			}
		}
	}

	return false
}
