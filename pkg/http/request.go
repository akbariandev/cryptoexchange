package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	MethodGET = http.MethodGet
)

func NewRequest(method, url string, body io.Reader, headers ...map[string][]string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type":   []string{"application/json"},
		"Content-Length": []string{strconv.FormatInt(req.ContentLength, 10)},
	}

	for _, h := range headers {
		for k, v := range h {
			req.Header.Set(k, v[0])
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("http response error code: %d", res.StatusCode))
	}

	return res, nil

}
