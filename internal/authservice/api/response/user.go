package response

import "github.com/google/uuid"

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"user_name"`
}

type UpdateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"user_name"`
}

type GetUserByIDResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"user_name"`
}

type GetAllUserResponse struct {
	Users []GetUserByIDResponse `json:"users"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
