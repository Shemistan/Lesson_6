package models

import "time"

type UserStatus int
type UserRole int

const Active UserStatus = 1
const Inactive UserStatus = 0

const Admin UserRole = 100
const Manager UserRole = 1

type User struct {
	Id               int64      `json:"id"`
	Login            string     `json:"-"`
	PasswordHash     string     `json:"-"`
	Name             string     `json:"name"`
	Surname          string     `json:"surname"`
	Status           UserStatus `json:"status"`
	Role             UserRole   `json:"role"`
	RegistrationDate time.Time  `json:"registration_date"`
	UpdateDate       time.Time  `json:"update_date"`
}
