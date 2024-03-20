package inn_http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/souzagabriel/glassblock-inn-service/pkg/util"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type DefaultInnHttpClient struct {
	apiKey     string
	shardUrl   string
	httpClient util.GenericHttpClient
}

func NewDefaultInnHttpClient(apiKey string, shardUrl string) (*DefaultInnHttpClient, error) {
	if !isApiKeyValid(apiKey) {
		return nil, errors.New("invalid apiKey, it should be a valid UUID")
	}

	validUrl, err := regexp.MatchString("^(https://)?.*\\..*$", shardUrl)

	if err != nil || validUrl == false {
		return nil, errors.New("invalid shardUrl, it should be a valid URL")
	}

	if !strings.HasPrefix(shardUrl, "https://") {
		shardUrl = "https://" + shardUrl
	}

	tr := &http.Transport{
		MaxIdleConns:       50,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: false,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   60 * time.Second,
	}

	return &DefaultInnHttpClient{
		apiKey:     apiKey,
		shardUrl:   shardUrl,
		httpClient: client,
	}, nil
}

func (c *DefaultInnHttpClient) Get(urlPath string) (*http.Response, error) {
	joinedUrl, err := url.JoinPath(c.shardUrl, urlPath)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Get(joinedUrl)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *DefaultInnHttpClient) Post(urlPath string, body map[string]interface{}) (*http.Response, error) {
	joinedUrl, err := url.JoinPath(c.shardUrl, urlPath)
	if err != nil {
		return nil, err
	}

	val, ok := body["apiKey"].(string)

	if !ok || !isApiKeyValid(val) {
		body["apiKey"] = c.apiKey
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return nil, errors.New("failed to marshal body map")
	}

	response, err := c.httpClient.Post(joinedUrl, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	return response, nil
}

func isApiKeyValid(apiKey string) bool {
	err := uuid.Validate(apiKey)
	if err != nil {
		return false
	}

	return true
}
