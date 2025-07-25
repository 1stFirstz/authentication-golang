package service

import "github.com/1stFirstz/authentication-golang/internal/authservice/repository"

type AuthService interface {
}

type authService struct {
	authRepository repository.UserRepository
}

func NewAuthService(authRepository repository.UserRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}
