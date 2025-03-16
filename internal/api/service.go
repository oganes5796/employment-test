package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/oganes5796/employment-test/internal/service"
)

type Implementation struct {
	todoService service.TodoService
}

func NewImplementation(todoService service.TodoService) *Implementation {
	return &Implementation{todoService: todoService}
}

func (i *Implementation) InitRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/todo", i.CreateTodo)
	app.Get("/todo", i.GetTodo)
	app.Put("/todo/:id", i.UpdateTodo)
	app.Delete("/todo/:id", i.DeleteTodo)

	app.Get("/swagger/*", swagger.HandlerDefault)

	return app
}
