package models

import "time"

type User struct {
	ID               int
	Login            string
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate time.Time
	UpdateDate       time.Time
}

type Statistics struct {
	DeleteUsersCount int
	UpdateCount      int
	GetUserCounts    int
	GetUsersCounts   int
}
