package util

import (
	"io"
	"net/http"
)

type GenericHttpClient interface {
	Get(url string) (resp *http.Response, err error)
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
}
