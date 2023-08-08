package accounts

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jbuget.fr/explore-golang/internal"
	"github.com/jbuget.fr/explore-golang/internal/core/domain"
)

type AccountsRepositoryPostgres struct {
	DB *internal.DB
}

func NewAccountsRepositoryPostgres(db *internal.DB) *AccountsRepositoryPostgres {
	return &AccountsRepositoryPostgres{DB: db}
}

func (repository *AccountsRepositoryPostgres) DeleteAccount(accountId int) {
	sqlStatement := `DELETE FROM accounts WHERE id = $1;`
	_, err := repository.DB.Client.Exec(sqlStatement, accountId)
	if err != nil {
		fmt.Println("No rows were deleted!")
		panic(err)
	}
}

func (repository *AccountsRepositoryPostgres) FindAccounts() []domain.Account {
	rows, _ := repository.DB.Client.Query("SELECT id, name, email, enabled FROM accounts;")
	defer rows.Close()

	var accounts []domain.Account

	for rows.Next() {
		var account domain.Account
		if err := rows.Scan(&account.Id, &account.Name, &account.Email, &account.Enabled); err != nil {
			return accounts
		}
		accounts = append(accounts, account)
	}

	return accounts
}

func (repository *AccountsRepositoryPostgres) GetAccountById(id int) domain.Account {
	var account domain.Account
	row := repository.DB.Client.QueryRow("SELECT id, name, email, enabled FROM accounts WHERE id=$1 LIMIT 1;", id)
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

func (repository *AccountsRepositoryPostgres) GetActiveAccountByEmail(email string) domain.AccountWithEncryptedPassword {
	var account domain.Account
	var accountWithEncryptedPassword domain.AccountWithEncryptedPassword

	sqlStatement := `
		SELECT id, created_at, name, email, password
		FROM accounts 
		WHERE email=$1
		AND enabled=true
		LIMIT 1;
	`
	row := repository.DB.Client.QueryRow(sqlStatement, email)
	switch err := row.Scan(&account.Id, &account.CreatedAt, &account.Name, &account.Email, &accountWithEncryptedPassword.EncryptedPassword); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		panic(err)
	case nil:
		account.Enabled = true
		accountWithEncryptedPassword.Account = account
		return accountWithEncryptedPassword
	default:
		panic(err)
	}
}

func (repository *AccountsRepositoryPostgres) InsertAccount(account domain.AccountWithEncryptedPassword) int {
	sqlStatement := `
		INSERT INTO accounts (name, email, password, enabled, created_at) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;`
	id := 0
	err := repository.DB.Client.QueryRow(sqlStatement, account.Account.Name, account.Account.Email, account.EncryptedPassword, account.Account.Enabled, account.Account.CreatedAt).Scan(&id)
	if err != nil {
		panic(err)
	}
	log.Println("New record ID is:", id)
	return id
}

func (repository *AccountsRepositoryPostgres) UpdateAccount(account domain.Account) domain.Account {
	panic("Not yet implemented")
}

func (repository *AccountsRepositoryPostgres) UpdatePassword(password string) domain.Account {
	panic("Not yet implemented")
}
