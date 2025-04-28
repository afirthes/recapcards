package main

import (
	"database/sql"
	"github.com/afirthes/recapcards/internal/db"
	"github.com/afirthes/recapcards/internal/env"
	"github.com/afirthes/recapcards/internal/store"
	"log"
)

func main() {
	log.Println("Starting server...")

	cfg := config{
		addr: env.GetString("SERVER_ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "30s"),
		},
	}

	database, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	// Closing database connection
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Println("Error closing database connection in defer section:", err)
		}
	}(database)

	storage := store.NewPostgresStorage(database)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	log.Printf("Server started at %s \n", cfg.addr)
	log.Fatal(app.Run(app.mount()))
}
