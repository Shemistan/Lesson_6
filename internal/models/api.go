package models

type AddRequest struct {
	AuthParams AuthData
	Date       UserDate
}

type AuthData struct {
	Login    string
	Password string
}

type UserDate struct {
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate string
	UpdateDate       string
}