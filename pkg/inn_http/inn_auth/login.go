package inn_auth

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/souzagabriel/glassblock-inn-service/pkg/inn_http"
	"github.com/souzagabriel/glassblock-inn-service/pkg/inn_http/inn_error"
)

type InnLoginBaseData struct {
	Password string `json:"-"`
	Otp      string `json:"otp"`
}

type InnLoginWithUserData struct {
	InnLoginBaseData
	User string `json:"user"`
}

type InnLoginWithEmailData struct {
	InnLoginBaseData
	Email string `json:"email"`
}

type InnLoginResponse struct {
	inn_error.InnErrorResponse
	inn_http.InnBaseResponse
	AppVerified   bool      `json:"appVerified"`
	ClusterShards []string  `json:"clusterShards"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"emailVerified"`
	KycStatus     string    `json:"kycStatus"`
	Pub           string    `json:"pub"`
	Roles         []string  `json:"roles"`
	Username      string    `json:"username"`
	UUID          uuid.UUID `json:"uuid"`
	Warning       string    `json:"warning"`
}

func Login[D InnLoginWithUserData | InnLoginWithEmailData](apiKey string, shardUrl string, data D) (*InnLoginResponse, error) {
	httpClient, err := inn_http.NewDefaultInnHttpClient(apiKey, shardUrl)
	if err != nil {
		return nil, err
	}

	innAuthClient := NewInnAuthClient(httpClient)

	switch any(data).(type) {
	case InnLoginWithUserData:
		if s, ok := any(data).(InnLoginWithUserData); ok {
			return innAuthClient.LoginWithUser(s)
		}
	case InnLoginWithEmailData:
		if s, ok := any(data).(InnLoginWithEmailData); ok {
			return innAuthClient.LoginWithEmail(s)
		}
	}

	return nil, errors.New("data entered is not valid")
}

func (i *InnAuthClient) LoginWithUser(data InnLoginWithUserData) (*InnLoginResponse, error) {
	body := make(map[string]interface{})

	body["user"] = data.User
	body["password"] = CreatePasswordHash(data.Password, data.User)
	body["otp"] = data.Otp

	response, err := i.innHttpClient.Post(LOGIN, body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fmt.Printf("Status code: %d\n", response.StatusCode)

	innBodyResponse := &InnLoginResponse{}
	err = ReadJSONBody(response, innBodyResponse)
	if err != nil {
		return nil, err
	}

	if innBodyResponse.Code != 0 {
		return nil, innBodyResponse
	}

	return innBodyResponse, nil
}

func (i *InnAuthClient) LoginWithEmail(data InnLoginWithEmailData) (*InnLoginResponse, error) {
	body := make(map[string]interface{})

	body["email"] = data.Email
	body["password"] = CreatePasswordHash(data.Password, data.Email)
	body["otp"] = data.Otp

	response, err := i.innHttpClient.Post(LOGIN, body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	innBodyResponse := &InnLoginResponse{}
	err = ReadJSONBody(response, innBodyResponse)
	if err != nil {
		return nil, err
	}

	if innBodyResponse.Code != 0 {
		return nil, innBodyResponse
	}

	return innBodyResponse, nil
}
