package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func NewApp() {
	client := newClient()
	repository := newRepository(client)
	service := newService(repository)
	handler := newHandler(service)
	app := fiber.New()
	newRouter(handler, app)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
