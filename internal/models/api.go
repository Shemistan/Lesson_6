package models

type AddRequest struct {
	AuthParams UserAuth
	Data       UserData
}

type UserAuth struct {
	Login    string
	Password string
}

type UserData struct {
	Name    string
	Surname string
}

type UpdateRequest struct {
	Id   int64
	Data UpdateUserData
}

type UpdateUserData struct {
	Name    string
	Surname string
	Status  string
	Role    string
}

type GetUserData struct {
	Id               int64
	Login            string
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate string
	UpdateDate       string
}
