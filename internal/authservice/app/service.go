package app

import (
	appService "github.com/1stFirstz/authentication-golang/internal/authservice/service"
)

type service struct {
	userService appService.UserService
}

func newService(repository *repository) *service {
	return &service{
		userService: appService.NewUserService(repository.userRepository),
	}
}
