package db

type Person struct {
	Name       string `dynamodbav:"p_name"`
	Age        int8   `dynamodbav:"p_age"`
	Profession string `dynamodbav:"p_prof"`
	Status     string `dynamodbav:"p_status"`
}

const (
	ATIVO   = "ATIVO"
	INATIVO = "INATIVO"
)

func NewPerson(name, profession string, age int8) Person {
	return Person{
		Name:       name,
		Profession: profession,
		Age:        age,
		Status:     ATIVO,
	}
}
