package dtos

type GetUserRequest struct {
	Id int `json:"id"`
}

type GetUserResponse []UserDto
