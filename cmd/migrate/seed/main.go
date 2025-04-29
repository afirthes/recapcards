package main

import (
	"database/sql"
	"github.com/afirthes/recapcards/internal/db"
	"github.com/afirthes/recapcards/internal/store"
	"log"
)

func main() {
	//addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New("postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable", 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Error closing database connection:", err)
		}
	}(conn)

	s := store.NewPostgresStorage(conn)

	db.Seed(s)
}
