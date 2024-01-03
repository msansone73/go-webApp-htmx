package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("goSansoneDB"))
	if err!=nil {
		log.Fatal(err.Error())
	}

	return db
}

