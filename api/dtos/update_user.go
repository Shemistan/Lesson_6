package dtos

type UpdateUserRequest struct {
	Id   int
	User UserDto
}

type UpdateUserResponse struct {
	Success bool
}
