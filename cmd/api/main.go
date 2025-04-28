package main

import (
	"log"
)

func main() {
	log.Println("Starting server...")

	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	log.Printf("Server started at %s \n", cfg.addr)
	log.Fatal(app.Run())
}
