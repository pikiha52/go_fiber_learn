package authRoutes

import (
	authHandler "go_fiber/internal/handler/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("auth")

	auth.Post("/signin", authHandler.Login)
}
