package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/naveenchander/GoKube/handlers"
)

var (
	// Release is a semantic version of current build
	Release = "unset"
	// PORT is a configurable via MakeFile
	PORT = "unset"
)

func main() {
	log.Printf(
		"Application Start - Release Version: %s",
		Release,
	)

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	router := handlers.Router(Release)

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}

	go func() {
		log.Fatalln(srv.ListenAndServe())
	}()

	log.Print("The service is ready to listen and serve on Port : " + PORT)

	killSignal := <-interupt

	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")

}
