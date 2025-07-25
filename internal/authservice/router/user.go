package router

import (
	"github.com/1stFirstz/authentication-golang/internal/authservice/handler"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	ApiRouter   fiber.Router
	UserHandler handler.UserHandler
}

func (r *UserRouter) NewUserRoutes() {
	userRouter := r.ApiRouter.Group("/users")
	userRouter.Post("/", r.UserHandler.CreateUser)
	userRouter.Get("/:id", r.UserHandler.GetUserByID)
	userRouter.Get("/", r.UserHandler.GetAllUser)
	userRouter.Put("/:id", r.UserHandler.UpdateUser)
	userRouter.Delete("/:id", r.UserHandler.DeleteUser)
}
