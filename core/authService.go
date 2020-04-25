package core

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
)

// ValidateAuthHeader ... Validate Authentication Header
func ValidateAuthHeader(header string) bool {
	hdr := strings.Split(header, " ")
	if len(hdr) != 2 {
		return false
	}

	decoded, err := base64.StdEncoding.DecodeString(hdr[1])
	if err != nil {
		return false
	}

	items := strings.Split(string(decoded), ":")

	if len(items) != 2 {
		return false
	}

	return validateClientAuthentication(items[0], HashString(items[1]))
}

// HashString ... Hash string
func HashString(data string) string {
	sh := sha256.Sum256([]byte(data))
	return string(sh[:])
}

func validateClientAuthentication(apiKey, secret string) bool {
	if apiKey == "" || secret == "" {
		return false
	}
	return dal.GetAuthHeader(apiKey, secret)
}

func AddClientCredentials(clientCredentials models.CustomerCredentials) {

}
