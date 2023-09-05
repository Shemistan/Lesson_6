package models

type SUser struct {
	Login            string
	PasswordHash     string
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate string
	UpdateDate       string
}

type IdGenerate struct {
	Id int32
}

type SAuth struct {
	Login        string
	PasswordHash string
}

// func (id *IdGenerate) GetId() int32 {
// 	id.IdGenerate++
// 	return id.IdGenerate
// }