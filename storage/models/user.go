package models

import "time"

// Statuses
const (
	StatusActive  string = "active"
	StatusDeleted        = "deleted"
)

// Roles
const (
	RoleAdmin     string = "Admin"
	RoleModerator string = "Moderator"
	RoleUser      string = "User"
)

type User struct {
	Id               int
	Login            string
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate time.Time
	UpdateDate       time.Time
}
