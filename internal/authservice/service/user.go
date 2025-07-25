package service

import (
	"context"
	"net/http"

	"github.com/1stFirstz/authentication-golang/internal/authservice/api/request"
	"github.com/1stFirstz/authentication-golang/internal/authservice/api/response"
	"github.com/1stFirstz/authentication-golang/internal/authservice/entity"
	"github.com/1stFirstz/authentication-golang/internal/authservice/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, request request.CreateUserRequest) (response.CreateUserResponse, int, error)
	GetUserByID(ctx context.Context, request request.GetUserByIDRequest) (response.GetUserByIDResponse, int, error)
	GetAllUser(ctx context.Context) ([]response.GetUserByIDResponse, int, error)
	UpdateUser(ctx context.Context, request request.UpdateUserRequest) (response.UpdateUserResponse, int, error)
	DeleteUser(ctx context.Context, request request.GetUserByIDRequest) (response.DeleteUserResponse, int, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(ctx context.Context, request request.CreateUserRequest) (response.CreateUserResponse, int, error) {
	user := &entity.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		UserName:  request.UserName,
		Password:  request.Password,
	}

	if err := s.userRepository.CreateUser(ctx, user); err != nil {
		return response.CreateUserResponse{}, http.StatusInternalServerError, err
	}

	res := response.CreateUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}

	return res, http.StatusOK, nil
}

func (s *userService) GetUserByID(ctx context.Context, request request.GetUserByIDRequest) (response.GetUserByIDResponse, int, error) {
	user, err := s.userRepository.GetUserByID(ctx, request.ID)
	if err != nil {
		return response.GetUserByIDResponse{}, http.StatusInternalServerError, err
	}

	res := response.GetUserByIDResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}

	return res, http.StatusOK, nil
}

func (s *userService) GetAllUser(ctx context.Context) ([]response.GetUserByIDResponse, int, error) {
	users, err := s.userRepository.GetAllUser(ctx)
	if err != nil {
		return []response.GetUserByIDResponse{}, http.StatusInternalServerError, err
	}

	res := []response.GetUserByIDResponse{}
	for _, user := range users {
		res = append(res, response.GetUserByIDResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			UserName:  user.UserName,
		})
	}

	return res, http.StatusOK, nil
}

func (s *userService) UpdateUser(ctx context.Context, request request.UpdateUserRequest) (response.UpdateUserResponse, int, error) {
	user, err := s.userRepository.GetUserByID(ctx, request.ID)
	if err != nil {
		return response.UpdateUserResponse{}, http.StatusInternalServerError, err
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.UserName = request.UserName
	user.Password = request.Password

	if err := s.userRepository.UpdateUser(ctx, user); err != nil {
		return response.UpdateUserResponse{}, http.StatusInternalServerError, err
	}

	res := response.UpdateUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}

	return res, http.StatusOK, nil
}

func (s *userService) DeleteUser(ctx context.Context, request request.GetUserByIDRequest) (response.DeleteUserResponse, int, error) {
	if err := s.userRepository.DeleteUser(ctx, request.ID); err != nil {
		return response.DeleteUserResponse{}, http.StatusInternalServerError, err
	}

	return response.DeleteUserResponse{
		Message: "User deleted successfully",
	}, http.StatusOK, nil
}
