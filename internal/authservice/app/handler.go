package app

import (
	appHandler "github.com/1stFirstz/authentication-golang/internal/authservice/handler"
)

type handler struct {
	userHandler appHandler.UserHandler
}

func newHandler(service *service) *handler {
	return &handler{
		userHandler: appHandler.NewUserHandler(service.userService),
	}
}
