package userRoutes

import (
	userHandler "go_fiber/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("user")

	user.Get("/", userHandler.GetUser)
	user.Post("/", userHandler.CreateUser)
}
