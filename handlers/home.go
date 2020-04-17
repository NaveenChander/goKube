package handlers

import (
	"fmt"
	"net/http"
)

func home(release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Welcome Home "+release)
	}
}
