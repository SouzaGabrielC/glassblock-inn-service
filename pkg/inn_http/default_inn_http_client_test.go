package inn_http

import (
	"encoding/json"
	"fmt"
	"io"
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
	defaultInnHttpClient, err := NewDefaultInnHttpClient(apiKey, shardUrl)

	if err != nil {
		t.Fatal(err)
	}

	body := make(map[string]interface{})
	body["uuid"] = "c9f7a257-5a9a-41b2-9bcf-52ba924c4298"

	resp, err := defaultInnHttpClient.Post("/core/wallet/assets", body)

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	respBodyJson := make(map[string]interface{})

	err = json.Unmarshal(respBody, &respBodyJson)

	if err != nil {
		t.Fatal(err)
	}

	for k, v := range respBodyJson {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}
