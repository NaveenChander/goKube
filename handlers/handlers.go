package handlers

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
	"github.com/naveenchander/GoKube/api"
)

// Router ... Define all external routes
func Router(release string) *mux.Router {

	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/home", api.Home(release)).Methods("GET")
	r.HandleFunc("/expMatch", api.ExpMatch(release)).Methods("POST")
	r.HandleFunc("/ClientCredentials", api.AddClientCredentials).Methods("POST")
	r.HandleFunc("/healthz", healthz).Methods("GET")
	r.HandleFunc("/readyz", readyz(isReady)).Methods("GET")
	return r
}
