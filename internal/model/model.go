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

type Stats struct {
	AddCounter      uint32
	UpdateCounter   uint32
	DeleteCounter   uint32
	GetUserCounter  uint32
	GetUsersCounter uint32
}
