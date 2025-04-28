package main

import (
	"github.com/afirthes/recapcards/internal/env"
	"log"
)

func main() {
	log.Println("Starting server...")

	cfg := config{
		addr: env.GetString("SERVER_ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	log.Printf("Server started at %s \n", cfg.addr)
	log.Fatal(app.Run(app.mount()))
}
