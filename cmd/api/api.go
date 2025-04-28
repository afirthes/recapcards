package main

import (
	"github.com/afirthes/recapcards/internal/store"
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
	config  config
	storage store.Storage
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Timeout on request context (ctx), that will signal
	// through ctx.Done() that the request has time out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)
	})

	return r
}

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))

	app.storage.Posts.Create(r.Context())

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
