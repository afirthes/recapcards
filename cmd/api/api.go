package main

import (
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

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", app.healthHandler)
	return mux
}

func (app *application) healthHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func (app *application) Run(mux *http.ServeMux) error {

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
