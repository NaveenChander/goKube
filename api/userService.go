package api

import (
	"net/http"
)

// AddUser ... Add a User
func AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
