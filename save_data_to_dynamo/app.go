package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/hugovallada/save_data_to_dynamo/src/db"
	"github.com/hugovallada/save_data_to_dynamo/src/resources"
)

type Table struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (tb Table) AddPerson(ctx context.Context, person Person) error {
	item, err := attributevalue.MarshalMap(person)
	if err != nil {
		return err
	}
	_, err = tb.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tb.TableName),
		Item:      item,
	})
	if err != nil {
		slog.Error("Couldn't add item to table")
	}
	return err
}

type Person struct {
	Name       string `json:"nome" dynamodbav:"tb_name"`
	Age        int8   `json:"idade" dynamodbav:"tb_age"`
	Profession string `json:"profissao" dynamodbav:"tb_profession"`
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch event.HTTPMethod {
	case "POST":
		person, err := resources.ConvertByteToStruct[Person]([]byte(event.Body))
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "Failure",
			}, err
		}
		tb := db.NewDynamoTable(&dynamodb.Client{}, "hlvl_db_person")
		if err := tb.AddPerson(ctx, person); err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "Failure",
			}, err
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 202,
			Body:       "Successfully created" + person.Name,
		}, err
	default:
		break
	}
	var person Person
	if err := json.Unmarshal([]byte(event.Body), &person); err != nil {
		slog.Error("Invalid payload")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Failed request",
		}, err
	}
	data, err := json.Marshal(person)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Failed request",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 202,
		Body:       string(data),
	}, nil
}

func createNewPerson(data []byte) (person Person, err error) {
	if err = json.Unmarshal(data, &person); err != nil {
		return person, err
	}
	return person, err
}

func main() {
	lambda.Start(handler)
}
