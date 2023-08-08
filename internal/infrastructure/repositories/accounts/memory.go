package accounts

import (
	"github.com/jbuget.fr/explore-golang/internal"
	"github.com/jbuget.fr/explore-golang/internal/core/domain"
)

type AccountsRepositoryMemory struct {
}

func NewAccountsRepositoryMemory(db *internal.DB) *AccountsRepositoryMemory {
	return &AccountsRepositoryMemory{}
}

func (repository *AccountsRepositoryMemory) DeleteAccount(accountId int) {
	panic("Not yet implemented")
}

func (repository *AccountsRepositoryMemory) FindAccounts() []domain.Account {
	panic("Not yet implemented")
}

func (repository *AccountsRepositoryMemory) GetAccountById(id int) domain.Account {
	panic("Not yet implemented")
}

func (repository *AccountsRepositoryMemory) GetActiveAccountByEmail(email string) domain.AccountWithEncryptedPassword {
	panic("Not yet implemented")
}

func (repository *AccountsRepositoryMemory) InsertAccount(account domain.AccountWithEncryptedPassword) int {
	panic("Not yet implemented")
}

func (repository *AccountsRepositoryMemory) UpdateAccount(account domain.Account) domain.Account {
	panic("Not yet implemented")
}

func (repository *AccountsRepositoryMemory) UpdatePassword(password string) domain.Account {
	panic("Not yet implemented")
}
