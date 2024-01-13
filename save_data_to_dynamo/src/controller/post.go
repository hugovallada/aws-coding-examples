package controller

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/hugovallada/save_data_to_dynamo/src/db"
	"github.com/hugovallada/save_data_to_dynamo/src/dto"
	"github.com/hugovallada/save_data_to_dynamo/src/resources"
	"github.com/hugovallada/save_data_to_dynamo/src/usecase"
)

func CreatePersonController(
	ctx context.Context, event events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	person, err := resources.ConvertByteToStruct[dto.Person]([]byte(event.Body))
	if err != nil {
		return ErrorResponse(400, "Failure"), err
	}
	createPerson := usecase.CreatePerson{Table: db.NewDynamoTable(&dynamodb.Client{}, "hlvl_db_person")}
	if err = createPerson.Execute(ctx, person); err != nil {
		return ErrorResponse(400, err.Error()), err
	}
	return SuccessResponse(201, "Successfully created a new person"), nil
}
