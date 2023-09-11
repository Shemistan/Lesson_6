package models

type Account struct {
	AuthData AuthData
	UserInfo UserInfo
}
type AuthData struct {
	Login    string
	Password string
}
type UserInfo struct {
	Name    string
	Surname string
	Active  bool
	Role    string
}
