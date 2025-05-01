package store

import (
	"database/sql"
	"log"
)

func CloseRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Println("Error closing rows:", err)
	}
}
