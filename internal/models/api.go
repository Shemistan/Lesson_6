package models

type AddUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type GetUserRequest struct {
	UserId int
}

type GetUsersRequest struct {
}

type UpdateUserRequest struct {
	Id       int
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type DeleteUserRequest struct {
	Id int
}
