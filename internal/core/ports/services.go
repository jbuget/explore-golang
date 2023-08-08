package ports

import "github.com/jbuget.fr/explore-golang/internal/core/domain"

type AccountsService interface {
	DeleteAccount(accountId int)
	FindAccounts() []domain.Account
	GetAccountById(id int) domain.Account
	GetActiveAccountByEmail(email string) domain.AccountWithEncryptedPassword
	InsertAccount(account domain.AccountWithEncryptedPassword) int
	UpdateAccount(account domain.Account) domain.Account
	UpdatePassword(password string) domain.Account
}
