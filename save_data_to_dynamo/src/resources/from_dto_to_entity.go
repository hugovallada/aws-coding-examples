package resources

import (
	"github.com/hugovallada/save_data_to_dynamo/src/db"
	"github.com/hugovallada/save_data_to_dynamo/src/dto"
)

func FromPersonDTOToPersonEntity(dto dto.Person) db.Person {
	return db.NewPerson(dto.Name, dto.Profession, dto.Age)
}
