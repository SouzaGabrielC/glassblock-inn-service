package util

import (
	"io"
	"net/http"
)

type GenericHttpClient interface {
	Get(url string) (resp *http.Response, err error)
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
}

type MockedGenericHttpClient struct {
	mockedResponse *http.Response
	mockedError    error
	Urls           []string
	ContentTypes   []string
	BodyContents   [][]byte
}

func NewMockedGenericHttpClient() *MockedGenericHttpClient {
	return &MockedGenericHttpClient{
		Urls:         make([]string, 0),
		ContentTypes: make([]string, 0),
		BodyContents: make([][]byte, 0),
	}
}

func (h *MockedGenericHttpClient) SetResponse(resp *http.Response) *MockedGenericHttpClient {
	h.mockedResponse = resp
	return h
}

func (h *MockedGenericHttpClient) SetError(err error) *MockedGenericHttpClient {
	h.mockedError = err
	return h
}

func (h *MockedGenericHttpClient) Get(url string) (resp *http.Response, err error) {
	h.Urls = append(h.Urls, url)
	return h.mockedResponse, h.mockedError
}

func (h *MockedGenericHttpClient) Clear() *MockedGenericHttpClient {
	h.Urls = make([]string, 0)
	h.ContentTypes = make([]string, 0)
	h.BodyContents = make([][]byte, 0)
	h.mockedError = nil
	h.mockedResponse = nil
	return h
}

func (h *MockedGenericHttpClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	h.Urls = append(h.Urls, url)
	h.ContentTypes = append(h.ContentTypes, contentType)

	bodyContent, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	h.BodyContents = append(h.BodyContents, bodyContent)

	return h.mockedResponse, h.mockedError
}
