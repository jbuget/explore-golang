package services

import (
	"github.com/jbuget.fr/explore-golang/internal/core/domain"
	"github.com/jbuget.fr/explore-golang/internal/core/ports"
)

type service struct {
	accountsRepository ports.AccountsRepository
}

func NewAccountsService(accountsRepository ports.AccountsRepository) *service {
	return &service{accountsRepository: accountsRepository}
}

func (srv *service) DeleteAccount(accountId int) {
	srv.accountsRepository.DeleteAccount(accountId)
}

func (srv *service) FindAccounts() []domain.Account {
	return srv.accountsRepository.FindAccounts()
}

func (srv *service) GetAccountById(id int) domain.Account {
	return srv.accountsRepository.GetAccountById(id)
}

func (srv *service) GetActiveAccountByEmail(email string) domain.AccountWithEncryptedPassword {
	return srv.accountsRepository.GetActiveAccountByEmail(email)
}

func (srv *service) InsertAccount(account domain.AccountWithEncryptedPassword) int {
	return srv.accountsRepository.InsertAccount(account)
}

func (srv *service) UpdateAccount(account domain.Account) domain.Account {
	return srv.accountsRepository.UpdateAccount(account)
}

func (srv *service) UpdatePassword(password string) domain.Account {
	return srv.accountsRepository.UpdatePassword(password)
}
