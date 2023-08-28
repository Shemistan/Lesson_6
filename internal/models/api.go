package models

import "time"

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	Status           string    `json:"status"`
	Role             string    `json:"role"`
	RegistrationDate time.Time `json:"registrationDate"`
	UpdateDate       time.Time `json:"updateDate"`
}
