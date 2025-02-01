package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

// We could have used chi.Mux as a return type
// But internally it implements http.Handler interface
// So we can return it as http.Handler to keep the implementation generic
// func (app *application) mount() *chi.Mux {}

func (app *application) mount() http.Handler {
	// This is the default way to create a new "ServeMux"
	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	// return mux

	// Can be installed using "go get -u github.com/go-chi/chi/v5"
	// Read for reference: https://github.com/go-chi/chi
	r := chi.NewRouter()

	// Middleware is a function that sits b/w request and response
	// This middleware is used to log the HTTP request
	r.Use(middleware.Logger)

	// A simple chi route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// Using nested routes to group the similar routes
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server started at port %s", app.config.addr)

	return srv.ListenAndServe()
}
