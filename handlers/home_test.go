package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	h := home("")
	h(w, nil)

	resp := w.Result()

	if returned, expected := resp.StatusCode, http.StatusOK; returned != expected {
		t.Errorf("Status code is wrong. returned: %d, expected: %d.", returned, expected)
	}

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	if returned, expected := string(greeting), "Welcome Home "; returned != expected {
		t.Errorf("The greeting is wrong. returned: %s, expected: %s.", returned, expected)
	}
}
