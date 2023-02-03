package v1

import (
	"bytes"
	"encoding/json"
	"gitlab.com/hotelian-company/challenge/pkg/logger"
	"net/http"
)

func NewRouter(handler *http.ServeMux) {
	NewRateRouter(handler)
	//implement other routes
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
