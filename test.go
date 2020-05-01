package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// startDate, _ := ptypes.TimestampProto(time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC))
	// strtDate, err := ptypes.Timestamp(startDate)

	// fmt.Println(strtDate)
	// if err != nil {
	// 	log.Fatal("Unable convert start Date")
	// }

	// key := "ZJBdn&6ckT&!2S(K4Qc^vJSyn"
	// secret := "N5$pEnkvZkd7Lp-)3B*aupY8r$t2#y3sCrBq2pGe"

	key := "testAPIKey"
	secret := "Abcd@1234"

	basic := fmt.Sprintf("%s:%s", key, secret)
	b := []byte(basic)

	result := base64.StdEncoding.EncodeToString(b)

	fmt.Println(result)
	fmt.Println("WkpCZG4mNmNrVCYhMlMoSzRRY152SlN5bjpONSRwRW5rdlprZDdMcC0pM0IqYXVwWThyJHQyI3kzc0NyQnEycEdl")
}
