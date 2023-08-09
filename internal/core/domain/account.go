package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Name      string
	Enabled   bool
}

func NewAccount(name string, email string, createdAt time.Time, updatedAt time.Time, enabled bool) Account {
	return Account{
		Name:      name,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Enabled:   enabled,
	}
}

type AccountWithEncryptedPassword struct {
	Account           Account
	EncryptedPassword string
}

func NewAccountWithEncryptedPassword(account Account, password string) AccountWithEncryptedPassword {
	var encryptedPassword string

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	} else {
		encryptedPassword = string(bytes)
	}

	accountWithPassword := AccountWithEncryptedPassword{
		Account:           account,
		EncryptedPassword: encryptedPassword,
	}
	return accountWithPassword
}
