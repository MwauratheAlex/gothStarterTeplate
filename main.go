package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/mwaurathealex/gothstarter/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Handle("/*", public())

	// router.Get expects a handleFunc that takes (w, r) and returns nothing
	// our function returns an error
	// handle Make is a decorator.
	// it takes in a our function, which returns an error,
	// prints the error then returns nothing,
	// making our function compatible with router.Get
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLoginIndex))

	listenAddress := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddress", listenAddress)

	http.ListenAndServe(listenAddress, router)
}
