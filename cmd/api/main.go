package main

import (
	"github.com/afirthes/recapcards/internal/env"
	"github.com/afirthes/recapcards/internal/store"
	"log"
)

func main() {
	log.Println("Starting server...")

	cfg := config{
		addr: env.GetString("SERVER_ADDR", ":8080"),
	}

	storage := store.NewPostgresStorage(nil)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	log.Printf("Server started at %s \n", cfg.addr)
	log.Fatal(app.Run(app.mount()))
}
