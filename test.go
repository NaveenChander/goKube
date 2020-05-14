package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	// startDate, _ := ptypes.TimestampProto(time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC))
	// strtDate, err := ptypes.Timestamp(startDate)

	// fmt.Println(strtDate)
	// if err != nil {
	// 	log.Fatal("Unable convert start Date")
	// }

	key := "JepWgR9hhxJTRS-E*v+D~BcP8"
	secret := "Qx+z$(T#kEDBSrRCnH9G+s&)L$t4Bgnct4&ka*aF"

	// key := "testAPIKey"
	// secret := "Abcd@1234"

	basic := fmt.Sprintf("%s:%s", key, secret)
	b := []byte(basic)

	result := base64.StdEncoding.EncodeToString(b)
	ret := ValidateAuthHeader("BASIC YyRoQyR3Y3ltJlpnajZKK1A1RXhkbXJNZTpQQnU2am5mZkVxPSM4KUA0YzlEdSFKUk52ck51UkByPThXUlZoY0A2")

	fmt.Println(ret)
	fmt.Println(result)
	fmt.Println("WkpCZG4mNmNrVCYhMlMoSzRRY152SlN5bjpONSRwRW5rdlprZDdMcC0pM0IqYXVwWThyJHQyI3kzc0NyQnEycEdl")

	/// YyRoQyR3Y3ltJlpnajZKK1A1RXhkbXJNZTpQQnU2am5mZkVxPSM4KUA0YzlEdSFKUk52ck51UkByPThXUlZoY0A2
}

// HashString ... Hash string
func HashString(data string) string {
	sh := sha256.Sum256([]byte(data))
	return strings.ToUpper(hex.EncodeToString(sh[:]))
}

func ValidateAuthHeader(header string) bool {

	fmt.Println("AuthHeader : " + header)

	hdr := strings.Split(header, " ")
	if len(hdr) != 2 {
		fmt.Println("Invalid Header")
		return false
	}

	decoded, err := base64.StdEncoding.DecodeString(hdr[1])
	if err != nil {
		fmt.Println("Error Decoding to Base64")
		return false
	}
	fmt.Println("Decoded header is : ")
	fmt.Println(decoded)

	items := strings.Split(string(decoded), ":")
	if len(items) != 2 {
		fmt.Println("Decoded Header not in correct format")
		return false
	}

	fmt.Println("API Key :" + items[0])
	fmt.Println("API secret :" + items[1])
	fmt.Println("API secret Hash:" + HashString(items[1]))
	return true
}
