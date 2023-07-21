package noteRoutes

import (
	noteHandler "go_fiber/internal/handler/note"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	note.Get("/", noteHandler.GetNotes)
	note.Post("/", noteHandler.CreateNote)
	note.Get("/:id", noteHandler.GetNote)
	note.Put("/:id", noteHandler.UpdateNote)
	note.Delete("/:id", noteHandler.DeleteNote)
}
