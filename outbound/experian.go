package outbound

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ExperianCall ... ExperianCall
type ExperianCall struct{}

// ExperianTestCall ... ExperianTestCall
type ExperianTestCall struct{}

// Dial ... Dial
func (ExperianCall) Dial(xmlToSend []byte, url, username, password string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlToSend))
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	basic := fmt.Sprintf("%s:%s", username, password)
	base64Basic := base64.StdEncoding.EncodeToString([]byte(basic))

	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	req.Header.Add("Basic", base64Basic)
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error receiving response from Experian : " + err.Error())
		log.Println(err)
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("Status error: " + string(res.StatusCode))
		return []byte{}, fmt.Errorf("Status error: %v", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
