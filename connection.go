package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	)

// get connection

func getConnection() *sql.DB {
	dsn := "postgres://user1:golang123@127.0.0.1:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
