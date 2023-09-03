package models

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

type Statistic struct {
	DeletedUsersCount int
	UpdateCount       int
	GetUserCount      int
	GetUsersCount     int
}
