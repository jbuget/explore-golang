package ports

import "github.com/jbuget.fr/explore-golang/internal/core/domain"

type AccountsService interface {
	DeleteAccount(accountId int)
	FindAccounts() []domain.Account
	GetAccountById(id int) domain.Account
	GetActiveAccountByEmail(email string) domain.AccountWithEncryptedPassword
	InsertAccount(name string, email string, password string) domain.Account
	UpdateAccount(id int, name string, email string) *domain.Account
	UpdatePassword(password string) domain.Account
}
