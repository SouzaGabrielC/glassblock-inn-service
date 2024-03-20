package util_test

import (
	"errors"
	"io"
	"net/http"
)

type MockedGenericHttpClient struct{}

func (h *MockedGenericHttpClient) SetGetResponse() {
}

func (h *MockedGenericHttpClient) Get(url string) (resp *http.Response, err error) {
	return nil, errors.New("")
}

func (h *MockedGenericHttpClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, errors.New("")
}
