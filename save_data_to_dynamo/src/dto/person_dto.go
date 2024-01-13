package dto

type Person struct {
	Name       string `json:"nome"`
	Age        int8   `json:"idade"`
	Profession string `json:"profissao"`
}
