package inn_http

import (
	"encoding/json"
	"errors"
	"github.com/souzagabriel/glassblock-inn-service/pkg/util"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestCreateNewDefaultInnHttpClient(t *testing.T) {
	validApiKey := "33a52d06-c5fe-4f36-bb50-ddaa3ef45060"
	validShardUrl := "https://test.devvio.com"

	t.Run("should return error when provide an invalid apiKey", func(t *testing.T) {
		_, err := NewDefaultInnHttpClient("abd123", validShardUrl)

		if err == nil {
			t.Error("Expected error on creating the DefaultInnHttpClient with invalid apiKey")
		}
	})

	t.Run("should return error when provide an invalid URL", func(t *testing.T) {
		_, err := NewDefaultInnHttpClient(validApiKey, "test")

		if err == nil {
			t.Error("Expected error on creating the DefaultInnHttpClient with invalid URL")
		}
	})

	t.Run("should not return error and create instance when all is correct", func(t *testing.T) {
		defaultInnHttpClient, err := NewDefaultInnHttpClient(validApiKey, validShardUrl)

		if err != nil {
			t.Error("Not expect error on creating the DefaultInnHttpClient")
		}

		if defaultInnHttpClient == nil {
			t.Error("Expect client not to be nil")
		}
	})

	t.Run("should add https:// if not provided", func(t *testing.T) {
		defaultInnHttpClient, _ := NewDefaultInnHttpClient(validApiKey, "test.devvio.com")

		if !strings.HasPrefix(defaultInnHttpClient.shardUrl, "https://") {
			t.Error("Expect shardUrl to have https://")
		}
	})
}

func TestDefaultInnHttpClientPost(t *testing.T) {
	apiKey := "d4d964fa-db59-4cd3-9c13-bf5b5c2711ff"
	shardUrl := "https://innovation.test.devvio.com"
	body := make(map[string]interface{})
	defaultResponse := &http.Response{
		StatusCode: http.StatusOK,
	}

	defaultInnHttpClient, err := NewDefaultInnHttpClient(apiKey, shardUrl)

	if err != nil {
		t.Fatal(err)
	}

	mockedHttpClient := util.NewMockedGenericHttpClient()
	clearMockedHttpClient := func() {
		mockedHttpClient.Clear()
	}

	defaultInnHttpClient.httpClient = mockedHttpClient

	t.Run("should pass through the error from HTTP client", func(t *testing.T) {
		defer clearMockedHttpClient()
		mockedHttpClient.SetError(errors.New("generic error"))

		_, err := defaultInnHttpClient.Post("/core/wallet/assets", body)

		if err == nil {
			t.Error("Error should not be nil when HTTP Client returns an error")
		}
	})

	t.Run("should call the HTTP client with the path + shard URL provided", func(t *testing.T) {
		defer clearMockedHttpClient()

		urlPath := "/core/wallet/assets"

		mockedHttpClient.SetResponse(defaultResponse)

		resp, _ := defaultInnHttpClient.Post(urlPath, body)

		joinedUrl, _ := url.JoinPath(shardUrl, urlPath)

		if joinedUrl != mockedHttpClient.Urls[0] {
			t.Error("URL in the HTTP client should be the same passed to method", joinedUrl, mockedHttpClient.Urls[0])
		}

		if resp.StatusCode != http.StatusOK {
			t.Error("Status code should be ok, same as the default response set", resp.StatusCode)
		}
	})

	t.Run("should inject api key into body", func(t *testing.T) {
		defer clearMockedHttpClient()
		urlPath := "/core/wallet/assets"
		mockedHttpClient.SetResponse(defaultResponse)

		defaultInnHttpClient.Post(urlPath, body)

		if len(mockedHttpClient.BodyContents) < 1 {
			t.Fatal("Body wasn't injected")
		}

		bodyClientFromRequest := make(map[string]interface{})

		err := json.Unmarshal(mockedHttpClient.BodyContents[0], &bodyClientFromRequest)
		if err != nil {
			t.Fatal("Error unmarshalling body from request", err)
		}

		apiKeyFromRequest, ok := bodyClientFromRequest["apiKey"]

		if !ok {
			t.Fatal("Api key wasn't injected into request body")
		}

		if apiKey != apiKeyFromRequest {
			t.Errorf("Api key from request doesn't match the api key from provided. [Request: %v, Provided: %v]", apiKeyFromRequest, apiKey)
		}
	})
}
