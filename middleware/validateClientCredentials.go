package middleware

import (
	"net/http"

	"github.com/naveenchander/GoKube/core"
)

// ValidateClientCredentials ... Validate Client Credentials
func ValidateClientCredentials(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !core.ValidateAuthHeader(authHeader) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
