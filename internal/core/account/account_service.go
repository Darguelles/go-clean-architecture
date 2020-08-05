package account

import (
	"sync"
)

/*
Take a look here :)
1.- This services uses a Repository to implement methods.
2.- You can have more than one implementation depending on your needs.
*/

type accountService struct {
	accountRepository AccountRepository
}

// NewAccountService generates a new AccountService instance
func NewAccountService(repository AccountRepository) AccountService {
	return &accountService{
		accountRepository: repository,
	}
}

func (a accountService) SaveAccount(account Account) {
	a.accountRepository.SaveAccount(account)
}

func (a accountService) GetAccounts() {
	// TODO implement this method as practice
	panic("implement me")
}

// Singleton
var onceService sync.Once
type singletonAccountService AccountService
var serviceInstance singletonAccountService

// DynamoRepositoryAdapterInstance creates a single instance of DynamoRepositoryAdapter
func AccountServiceInstance() AccountService {
	onceService.Do(func() {
		serviceInstance = NewAccountService(DynamoAccountRepositoryInstance())
	})
	return serviceInstance
}


