package outbound

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// IExperian ... IExperian
type IExperian interface {
	Dial(xmlToSend []byte) ([]byte, error)
	SetExperianDetails(url, username, password string)
}

// ExperianCall ... ExperianCall
type ExperianCall struct {
	url, username, password string
}

// ExperianTestCall ... ExperianTestCall
type ExperianTestCall struct {
	url, username, password string
}

// SetExperianDetails ... SetExperianDetails
func (c ExperianTestCall) SetExperianDetails(url, username, password string) {
	c.url = url
	c.username = username
	c.password = password
	return
}

// Dial ... Dial
func (c ExperianTestCall) Dial(xmlToSend []byte) ([]byte, error) {
	return xmlToSend, nil
}

// SetExperianDetails ... SetExperianDetails
func (c ExperianCall) SetExperianDetails(url, username, password string) {
	c.url = url
	c.username = username
	c.password = password
	return
}

// Dial ... Dial
func (c ExperianCall) Dial(xmlToSend []byte) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(xmlToSend))
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	basic := fmt.Sprintf("%s:%s", c.username, c.password)
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

	log.Println("Data response from Experian ---> " + string(data))
	return data, nil
}
