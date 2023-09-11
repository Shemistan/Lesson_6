package models

type User struct {
	Login            string
	Password         string
	Name             string
	Surname          string
	Active           bool
	Role             string
	RegistrationDate string
	UpdateDate       string
}
