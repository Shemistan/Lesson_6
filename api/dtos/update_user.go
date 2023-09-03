package dtos

type UpdateUserRequest struct {
	id   int
	user UserDto
}

type UpdateUserResponse struct {
	success bool
}
