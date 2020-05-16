package core

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
	"strings"

	"github.com/naveenchander/GoKube/dal"
)

// ValidateAuthHeader ... Validate Authentication Header
func ValidateAuthHeader(header string, serviceName string) bool {

	log.Println("AuthHeader : " + header)

	hdr := strings.Split(header, " ")
	if len(hdr) != 2 {
		log.Println("Invalid Header")
		return false
	}

	decoded, err := base64.StdEncoding.DecodeString(hdr[1])
	if err != nil {
		log.Println("Error Decoding to Base64")
		return false
	}
	log.Println("Decoded header is : ")
	log.Println(decoded)

	items := strings.Split(string(decoded), ":")
	if len(items) != 2 {
		log.Println("Decoded Header not in correct format")
		return false
	}

	log.Println("API Key :" + items[0])
	log.Println("API secret :" + items[1])
	log.Println("API secret Hash:" + HashString(items[1]))

	return validateClientAuthentication(items[0], HashString(items[1]), serviceName)
}

// HashString ... Hash string
func HashString(data string) string {
	sh := sha256.Sum256([]byte(data))
	return strings.ToUpper(hex.EncodeToString(sh[:]))
}

func validateClientAuthentication(apiKey, secret string, serviceName string) bool {
	if apiKey == "" || secret == "" {
		return false
	}
	return dal.GetAuthHeader(apiKey, secret) && dal.IsAuthorized(apiKey, serviceName)
}
