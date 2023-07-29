package database

import (
	"database/sql"
	"log"
)

type DB struct {
	Client *sql.DB
	URL    string
}

func Connect(url string) (*DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &DB{
		Client: db,
		URL:    url,
	}, nil
}
