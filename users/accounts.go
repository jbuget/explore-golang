package users

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jbuget.fr/explore-golang/database"
)

type Account struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Name      string
	Enabled   bool
}

type AccountWithEncryptedPassword struct {
	Account           Account
	EncryptedPassword string
}

type AccountRepository struct {
	DB *database.DB
}

func CreateAccount(name string, email string, password string) AccountWithEncryptedPassword {
	account := Account{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		Enabled:   true,
	}
	accountWithPassword := AccountWithEncryptedPassword{
		Account:           account,
		EncryptedPassword: password,
	}
	return accountWithPassword
}

func (repository *AccountRepository) InsertAccount(account AccountWithEncryptedPassword) (int) {
	sqlStatement := `
INSERT INTO accounts (name, email, password, enabled) 
VALUES ($1, $2, $3, $4)
RETURNING id`
	id := 0
	err := repository.DB.Client.QueryRow(sqlStatement, account.Account.Name, account.Account.Email, account.EncryptedPassword, account.Account.Enabled).Scan(&id)
	if err != nil {
		panic(err)
	}
	log.Println("New record ID is:", id)
	return id
}

func (repository *AccountRepository) UpdateAccount(account Account) Account {
	panic("Not yet implemented")
}

func (repository *AccountRepository) UpdatePassword(password string) Account {
	panic("Not yet implemented")
}

func (repository *AccountRepository) GetAccount() Account {
	var account Account
	row := repository.DB.Client.QueryRow("SELECT id, name, email, enabled FROM accounts WHERE email='david@example.org' LIMIT 1")
	switch err := row.Scan(&account.Id, &account.Name, &account.Email, &account.Enabled); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		panic(err)
	case nil:
		return account
	default:
		panic(err)
	}
}

func (repository *AccountRepository) FindAccounts() []Account {
	rows, _ := repository.DB.Client.Query("SELECT id, name, email, enabled FROM accounts")
	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.Id, &account.Name, &account.Email, &account.Enabled); err != nil {
			return accounts
		}
		accounts = append(accounts, account)
	}

	return accounts
}
