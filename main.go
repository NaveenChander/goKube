package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/naveenchander/GoKube/configuration"
	"github.com/naveenchander/GoKube/core"
	"github.com/naveenchander/GoKube/handlers"
)

func main() {
	log.Printf(
		"Application Start - Release Version: %s",
		configuration.Release,
	)

	core.InitFlakeIDGenerator()

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	router := handlers.Router(configuration.Release)

	srv := &http.Server{
		Addr:    ":" + configuration.PORT,
		Handler: router,
	}

	go func() {
		log.Fatalln(srv.ListenAndServe())
	}()

	log.Print("The service is ready to listen and serve on Port : " + configuration.PORT)

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
