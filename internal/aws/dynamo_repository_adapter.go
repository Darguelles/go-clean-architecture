package aws

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"sync"
)

/*
Take a look here :)
1.- Here we use the AWS SDK methods and perform IO operations
2.- Here data comes and is transformed into AWS SDK payload on request
3.- Responses are returned into application objects, no AWS SDK objects.
*/

// DynamoRepositoryAdapter defines the contract for this Repository
type DynamoRepositoryAdapter interface {
	Save(data interface{}, tableName string) error
}

type dynamoRepositoryAdapter struct {
}

// Save stores an element on DynamoDB, returns an error only when something goes wrong
func (r *dynamoRepositoryAdapter) Save(data interface{}, tableName string) error {
	sess := GetSession()
	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		return errors.New("Error! Put a message explaining the error")
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		return errors.New("Error! Put a message explaining the error")
	}
	return err
}

// NewDynamoRepositoryAdapter creates a new instance of DynamoRepositoryAdapter each time is called
func NewDynamoRepositoryAdapter() DynamoRepositoryAdapter {
	return &dynamoRepositoryAdapter{}
}

// Singleton
var once sync.Once
type singletonDynamoRepositoryAdapter DynamoRepositoryAdapter

var instance singletonDynamoRepositoryAdapter

// DynamoRepositoryAdapterInstance creates a single instance of DynamoRepositoryAdapter
func DynamoRepositoryAdapterInstance() DynamoRepositoryAdapter {
	once.Do(func() {
		instance = NewDynamoRepositoryAdapter()
	})
	return instance
}
