package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	app1 := newApp()

	router := chi.NewRouter()
	router.Get("/data", app1.GetData)

	port := ":9090"
	srv := http.Server{
		Addr:        port,
		Handler:     router,
		IdleTimeout: 120 * time.Second,
	}
	log.Printf("SERVER STARTING ON PORT:%v \n\n", port)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Printf("ERROR STARTING SERVER: %v", err)
			os.Exit(1)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Printf("Closing now, We've gotten signal: %v", sig)

	ctx := context.TODO()
	srv.Shutdown(ctx)
}
