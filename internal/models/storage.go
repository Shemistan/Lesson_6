package models

type User struct {
	Name             string
	Surname          string
	Login            string
	HashPassword     string
	Status           string
	Role             string
	RegistrationDate string
	UpdateDate       string
}