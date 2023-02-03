package v1

import (
	"bytes"
	"encoding/json"
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
		//TODO LOGGER
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write(nil); err != nil {
			defer func() {
				if r := recover(); r != nil {
					//TODO LOGGER
				}

				panic(err)
			}()
			return
		}
	}

	if _, err := w.Write(reqBodyBytes.Bytes()); err != nil {
		//TODO LOGGER
		defer func() {
			if r := recover(); r != nil {
				//TODO LOGGER
			}

			panic(err)
		}()
	}
}
