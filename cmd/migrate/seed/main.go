package main

import (
	"github.com/afirthes/recapcards/internal/db"
	"github.com/afirthes/recapcards/internal/store"
	"log"
)

func main() {
	conn, err := db.New("postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable", 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewPostgresStorage(conn)

	db.Seed(store, conn)
}
