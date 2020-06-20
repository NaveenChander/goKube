package outbound

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// IIDMatch ... IIDMatch
type IIDMatch interface {
	Dial(xmlToSend []byte) ([]byte, error)
	SetIDMatchDetails(url, username, password string)
}

// IDMatchCall ... IDMatchCall
type IDMatchCall struct {
	url, username, password string
}

// IDMatchTestCall ... IDMatchTestCall
type IDMatchTestCall struct {
	url, username, password string
}

// SetIDMatchDetails ... SetIDMatchDetails
func (c IDMatchTestCall) SetIDMatchDetails(url, username, password string) {
	c.url = url
	c.username = username
	c.password = password
	return
}

// Dial ... Dial
func (c IDMatchTestCall) Dial(xmlToSend []byte) ([]byte, error) {
	return xmlToSend, nil
}

// SetIDMatchDetails ... SetIDMatchDetails
func (c IDMatchCall) SetIDMatchDetails(url, username, password string) {
	c.url = url
	c.username = username
	c.password = password
	return
}

// Dial ... Dial
func (c IDMatchCall) Dial(xmlToSend []byte) ([]byte, error) {

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
		log.Println("Error receiving response from IDMatch : " + err.Error())
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

	log.Println("Data response from IDMatch ---> " + string(data))
	return data, nil
}
