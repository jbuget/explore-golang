package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Account struct {
	Id       int
	Email    string
	Name     string
	Password string
	Enabled  bool
}

func GetAccount() Account {
	databaseUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var email string
	var name string
	var password string
	var enabled bool

	row := db.QueryRow("SELECT * FROM accounts WHERE email='clara@example.org' LIMIT 1")
	switch err := row.Scan(&id, &name, &email, &password, &enabled); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		panic(err)
	case nil:
		account := Account{
			Id:       id,
			Email:    email,
			Name:     name,
			Password: password,
			Enabled:  enabled,
		}
		return account
	default:
		panic(err)
	}
}

func FindAccounts() []Account {
	alice := Account{
		Id:       1,
		Email:    "alice@example.com",
		Name:     "Alice",
		Password: "abcd1234",
		Enabled:  true,
	}
	bob := Account{
		Id:       2,
		Email:    "bob@example.com",
		Name:     "Bob",
		Password: "abcd1234",
		Enabled:  true,
	}

	accounts := []Account{alice, bob}
	return accounts
}
