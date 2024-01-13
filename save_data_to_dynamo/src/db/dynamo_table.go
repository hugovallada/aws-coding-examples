package db

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoTable struct {
	DynamoClient *dynamodb.Client
	TableName    string
}

func (dt DynamoTable) InsertPerson(ctx context.Context, person Person) error {
	item, err := attributevalue.MarshalMap(person)
	if err != nil {
		return err
	}
	_, err = dt.DynamoClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(dt.TableName),
		Item:      item,
	})
	if err != nil {
		slog.Error("Não foi possível persistir no banco")
	}
	return err
}

func NewDynamoTable(dynamoClient *dynamodb.Client, tableName string) DynamoTable {
	return DynamoTable{
		DynamoClient: dynamoClient,
		TableName:    tableName,
	}
}
