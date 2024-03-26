package inn_auth

import "github.com/souzagabriel/glassblock-inn-service/pkg/inn_http"

type InnAuthClient struct {
	innHttpClient *inn_http.DefaultInnHttpClient
}

func NewInnAuthClient(httpClient *inn_http.DefaultInnHttpClient) *InnAuthClient {
	return &InnAuthClient{
		httpClient,
	}
}
