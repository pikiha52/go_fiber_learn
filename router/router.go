package router

import (
	authRoutes "go_fiber/internal/routes/auth"
	noteRoutes "go_fiber/internal/routes/note"
	userRoutes "go_fiber/internal/routes/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	noteRoutes.SetupNoteRoutes(api)
	userRoutes.SetupUserRoutes(api)
	authRoutes.SetupAuthRoutes(api)

}
