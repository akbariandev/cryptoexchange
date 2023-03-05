package v1

import (
	"encoding/json"
	"github.com/akbarian.dev/cryptoexchange/internal/core/rate"
	"github.com/akbarian.dev/cryptoexchange/pkg/logger"
	"net/http"
)

type rateRoutes struct{}

type getRateRequest struct {
	Currencies []string `json:"currencies"`
	To         string   `json:"to"`
}

type getRateResponse struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

func NewRateRouter() *rateRoutes {
	return &rateRoutes{}
}

func (r *rateRoutes) getRate(w http.ResponseWriter, req *http.Request) {
	body := getRateRequest{}
	ctx := req.Context()
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		logger.Logger.Error(err.Error())
		response(w, http.StatusInternalServerError, nil)
	}

	rates, err := rate.GetCurrenciesRate(ctx, body.Currencies, body.To)
	if err != nil {
		logger.Logger.Error(err.Error())
		response(w, http.StatusInternalServerError, err.Error())
	}

	resp := []getRateResponse{}
	for i, c := range body.Currencies {
		if rates[i] != 0 {
			resp = append(resp, getRateResponse{
				From: c,
				To:   body.To,
				Rate: rates[i],
			})
		}
	}

	response(w, http.StatusOK, resp)
}
