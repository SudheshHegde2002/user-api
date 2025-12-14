package routes

import (
	"github.com/gofiber/fiber/v2"

	"user-api/internal/handler"
)

func RegisterUserRoutes(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.CreateUser)
	app.Get("/users/:id", h.GetUserByID)
	app.Get("/users", h.ListUsers)
}
