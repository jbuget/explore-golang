package ports

import "github.com/jbuget.fr/explore-golang/internal/core/domain"

type AccountsRepository interface {
	DeleteAccount(accountId int)
	FindAccounts() []domain.Account
	GetAccountById(id int) domain.Account
	GetActiveAccountByEmail(email string) domain.AccountWithEncryptedPassword
	InsertAccount(account domain.AccountWithEncryptedPassword) int
	UpdateAccount(account *domain.Account)
	UpdatePassword(password string) domain.Account
}
