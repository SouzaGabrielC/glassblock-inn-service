package inn_auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func Test_CreatePasswordHash(t *testing.T) {
	email := "test@test.com"
	password := "Test123"

	expectedHashedPass := "820500fccf2e5a00ed793e005869c96f41db456a06e49411152a58c11b6ff2ea"

	hashedPass := CreatePasswordHash(email, password)

	if expectedHashedPass != hashedPass {
		t.Fatal("Expected hashed pass is not equal to actual hashed pass", expectedHashedPass, hashedPass)
	}
}

func Test_ReadInnBody(t *testing.T) {
	type TestResponse struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	bodyTestResponse := &TestResponse{
		Id:   "123",
		Name: "Test",
	}

	jsonBytes, _ := json.Marshal(bodyTestResponse)
	reader := io.NopCloser(bytes.NewReader(jsonBytes))

	headers := http.Header{}

	headers.Set("Content-Type", "application/json")

	response := &http.Response{
		Body:   reader,
		Header: headers,
	}

	var data *TestResponse = &TestResponse{}

	err := ReadJSONBody(response, data)
	if err != nil {
		t.Fatal("Error reading body", err)
	}

	if data.Id != bodyTestResponse.Id {
		t.Fatal("Read response Id is different than provided bodyTestResponse Id.", data.Id, bodyTestResponse.Id)
	}

	if data.Name != bodyTestResponse.Name {
		t.Fatal("Read response Name is different than provided bodyTestResponse Name.", data.Name, bodyTestResponse.Name)
	}
}
