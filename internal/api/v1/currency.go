package v1

import (
	"gitlab.com/hotelian-company/challenge/config"
	"gitlab.com/hotelian-company/challenge/pkg/logger"
	"net/http"
)

type currencyRoutes struct{}

type getCurrencyListResponse struct {
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
}

func NewCurrencyRouter() *currencyRoutes {
	return &currencyRoutes{}
}

func (r *currencyRoutes) getCurrencies(w http.ResponseWriter, req *http.Request) {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Logger.Error(err.Error())
		response(w, http.StatusInternalServerError, err.Error())
	}

	resp := []getCurrencyListResponse{}
	for currencyName, currencyConfig := range cfg.Currencies {
		resp = append(resp, getCurrencyListResponse{
			Enable: currencyConfig.Enable,
			Name:   currencyName,
		})
	}

	response(w, http.StatusOK, resp)
}
