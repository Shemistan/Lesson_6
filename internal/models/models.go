package models

type User struct {
	Id               int64
	Login            string
	Password         string
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate string
	UpdateDate       string
}

type Statistic struct {
	UsersCount          int64
	DeletedUsersAccount int64
	UpdateCount         int64
	GetUserCount        int64
	GetUsersCount       int64
}
