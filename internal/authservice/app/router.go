package app

import (
	appRouter "github.com/1stFirstz/authentication-golang/internal/authservice/router"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	userRouter appRouter.UserRouter
}

func newRouter(handler *handler, app *fiber.App) {
	initApiRouter := app.Group("/api/v1")
	router := &router{
		userRouter: appRouter.UserRouter{
			ApiRouter:   initApiRouter,
			UserHandler: handler.userHandler,
		},
	}

	router.userRouter.NewUserRoutes()
}
