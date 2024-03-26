package inn_auth

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func CreatePasswordHash(password string, usernameOrEmail string) string {
	stringToHash := password + strings.ToLower(usernameOrEmail)

	hash := sha256.New()
	hash.Write([]byte(stringToHash))

	return hex.EncodeToString(hash.Sum(nil))
}

func ReadJSONBody[B any](response *http.Response, data *B) error {
	if response.Header.Get("Content-Type") != "application/json" {
		return errors.New("content type is different than 'application/json'")
	}

	reader := bufio.NewReader(response.Body)
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		return err
	}

	return nil
}
