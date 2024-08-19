package accrual

import "net/http"

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type urlBuilder interface {
	Build(endpoint string, params map[string]string) string
}
