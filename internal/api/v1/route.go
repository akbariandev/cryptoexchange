package v1

import (
	"bytes"
	"encoding/json"
	"gitlab.com/hotelian-company/challenge/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

type routeHandler struct {
	rateRouteHandler     *rateRoutes
	currencyRouteHandler *currencyRoutes
}

func (r *routeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		switch req.URL.String() {
		case "/v1/currencies":
			r.currencyRouteHandler.getCurrencies(w, req)
		default:
			http.NotFound(w, req)
		}
	case http.MethodPost:
		switch req.URL.String() {
		case "/v1/getRate":
			r.rateRouteHandler.getRate(w, req)
		default:
			http.NotFound(w, req)
		}
	default:
		logger.Logger.Info("unhandled route", zap.String("url", req.URL.String()))
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(nil)
	}
}

func NewRouter(handler *http.ServeMux) {
	handler.Handle("/", &routeHandler{
		rateRouteHandler:     NewRateRouter(),
		currencyRouteHandler: NewCurrencyRouter(),
	})
}

func response(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	reqBodyBytes := new(bytes.Buffer)
	if err := json.NewEncoder(reqBodyBytes).Encode(data); err != nil {
		logger.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write(nil); err != nil {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error(err.Error())
				}

				panic(err)
			}()
			return
		}
	}

	if _, err := w.Write(reqBodyBytes.Bytes()); err != nil {
		logger.Logger.Error(err.Error())
		defer func() {
			if r := recover(); r != nil {
				logger.Logger.Error(err.Error())
			}

			panic(err)
		}()
	}
}
