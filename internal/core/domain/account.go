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

type AccountWithEncryptedPassword struct {
	Account           Account
	EncryptedPassword string
}

func NewAccountWithEncryptedPassword(name string, email string, password string) AccountWithEncryptedPassword {
	account := Account{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		Enabled:   true,
	}
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
