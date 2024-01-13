package dto

import "errors"

type Person struct {
	Name       string `json:"nome"`
	Age        int8   `json:"idade"`
	Profession string `json:"profissao"`
}

func (p Person) Validate() error {
	if p.Name == "" || len(p.Name) < 3 {
		return errors.New("invalid name")
	} else if p.Age < 0 {
		return errors.New("invalid age")
	} else if p.Profession == "" || len(p.Profession) < 2 {
		return errors.New("invalid profession")
	} else {
		return nil
	}
}
