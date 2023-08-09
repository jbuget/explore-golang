package services

import (
	"time"

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

func (srv *service) InsertAccount(name string, email string, password string) domain.Account {
	account := domain.NewAccount(name, email, time.Now(), time.Time{}, true)
	accountWithPassword := domain.NewAccountWithEncryptedPassword(account, password)
	id := srv.accountsRepository.InsertAccount(accountWithPassword)
	account.Id = id
	return account
}

func (srv *service) UpdateAccount(id int, name string, email string) *domain.Account {
	account := srv.accountsRepository.GetAccountById(id)
	if (account == domain.Account{}) {
		return nil
	}

	if name != "" {
		account.Name = name
	}
	if email != "" {
		account.Email = email
	}

	srv.accountsRepository.UpdateAccount(&account)
	return &account
}

func (srv *service) UpdatePassword(password string) domain.Account {
	return srv.accountsRepository.UpdatePassword(password)
}
