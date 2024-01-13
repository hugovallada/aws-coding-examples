package usecase

import (
	"context"

	"github.com/hugovallada/save_data_to_dynamo/src/db"
	"github.com/hugovallada/save_data_to_dynamo/src/dto"
	"github.com/hugovallada/save_data_to_dynamo/src/resources"
)


type CreatePerson struct {
	Table db.DynamoTable
}


func (cp CreatePerson) Execute(ctx context.Context, person dto.Person) error {
	if err := person.Validate(); err != nil {
		return err
	}
	return cp.Table.InsertPerson(ctx, resources.FromPersonDTOToPersonEntity(person))
}