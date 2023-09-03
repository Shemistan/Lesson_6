package models

type SUser struct {
	Login            string
	PasswordHash     string
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate string
	UpdateDate       string
}

// type JUser struct{
// 	SUser
// }

// type JAuth struct{
// 	SAuth
// }

type IdGenerate struct {
	IdGenerate int32
}

type SAuth struct {
	Login        string
	PasswordHash string
}

// type StatsUser struct{
// 	DeleteUsersCount int32
// 	UpdateCount int32
// 	GetUserCount int32
// 	GetUsersCount int32
// 	GetAuthClick int32
// }

func (id *IdGenerate) GetId() int32 {
	id.IdGenerate++
	return id.IdGenerate
}