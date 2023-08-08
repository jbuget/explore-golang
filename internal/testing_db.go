package internal

import (
	"log"
	"os"
)

var TestingDB DB

func GetTestingDB() (*DB) {
	if TestingDB == (DB{}) {
		databaseUrl := "postgres://test:test@localhost:15434/test?sslmode=disable"
		db, err := Connect(databaseUrl)
		if err != nil {
			log.Printf("error: %v\n", err)
			os.Exit(1)
		}
		log.Println("Database connected")
		TestingDB = *db
	}
	return &TestingDB
}
