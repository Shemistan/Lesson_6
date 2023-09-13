package model

type User struct {
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate string
	UpdateDate       string
	Login            string
	HashPassword     uint32
}
