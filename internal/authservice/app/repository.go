package app

import (
	appRepository "github.com/1stFirstz/authentication-golang/internal/authservice/repository"
)

type repository struct {
	userRepository appRepository.UserRepository
}

func newRepository(client *client) *repository {
	return &repository{
		userRepository: appRepository.NewUserRepository(client.db),
	}
}
