package models

// You use User model to save user information to DB
type User struct {
	Id        int64
	Login     string
	Firstname string
	Lastname  string
	Status string
	HashedPassword string
	RegistrationDate string
	UpdatedDate string
}

