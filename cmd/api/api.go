package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type config struct {
	addr string
}

type application struct {
	config config
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)
	})

	return r
}

func (app *application) healthHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func (app *application) Run(mux http.Handler) error {

	if app.config.addr == "" {
		log.Fatalf("FATAL: Server address is not set")
	}

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
