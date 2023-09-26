package models

type Request struct {
	AuthParams AuthParams
}

type AuthParams struct {
	Login string
	Password string
}

type UserUpdateRequest struct {
	Firstname string
	Lastname string
}
