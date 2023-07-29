package users

import (
	"database/sql"
	"fmt"

	"github.com/jbuget.fr/explore-golang/database"
)

type Account struct {
	Id       int
	Email    string
	Name     string
	Password string
	Enabled  bool
}

type AccountRepository struct {
	DB *database.DB
}

func (repository *AccountRepository) GetAccount() Account {
	var id int
	var email string
	var name string
	var password string
	var enabled bool

	row := repository.DB.Client.QueryRow("SELECT * FROM accounts WHERE email='david@example.org' LIMIT 1")
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

func (repository *AccountRepository) FindAccounts() []Account {
	rows, _ := repository.DB.Client.Query("SELECT * FROM accounts")
	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.Id, &account.Name, &account.Email, &account.Password, &account.Enabled); err != nil {
			return accounts
		}
		accounts = append(accounts, account)
	}

	return accounts
}
