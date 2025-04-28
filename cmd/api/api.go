package main

import (
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
	mux.HandleFunc("/api/v1/cards", app.cardsHandler)
	return mux
}

func (app *application) cardsHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		return
	}
}

func (app *application) Run() error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.mount(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return srv.ListenAndServe()
}
