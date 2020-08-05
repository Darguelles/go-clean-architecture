package account

import (
	"github.com/sirupsen/logrus"
	"go-clean-architecture/internal/aws"
	"sync"
)

/*
Take a look here :)
1.- No AWS SDK packages in this layer
2.- Repository for the AWS DynamoDB service based on AccountRepository interface
*/

type dynamoAccountRepository struct {
	dynamoRepositoryAdapter aws.DynamoRepositoryAdapter
}

// NewDynamoAccountRepository generates a new Repository using AccountRepository
func NewDynamoAccountRepository(adapter aws.DynamoRepositoryAdapter) AccountRepository {
	return &dynamoAccountRepository{
		dynamoRepositoryAdapter: adapter,
	}
}

func (a dynamoAccountRepository) SaveAccount(account Account) {
	tableName := "AccountTable"
	err := a.dynamoRepositoryAdapter.Save(account, tableName)
	if err != nil {
		logrus.Error("Error, put some self-explanatory message here")
	}
}

func (a dynamoAccountRepository) GetAccounts() {
	// TODO implement this method as practice
	panic("implement me")
}

// Singleton
var onceRepo sync.Once
type singletonDynamoAccountRepository AccountRepository

var repositoryInstance singletonDynamoAccountRepository

// DynamoRepositoryAdapterInstance creates a single instance of DynamoRepositoryAdapter
func DynamoAccountRepositoryInstance() AccountRepository {
	onceRepo.Do(func() {
		repositoryInstance = NewDynamoAccountRepository(aws.DynamoRepositoryAdapterInstance())
	})
	return repositoryInstance
}

