package userHandler

import (
	"go_fiber/database"
	"go_fiber/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No users present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "users Found", "data": users})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user.ID = uuid.New()

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	err = db.Create(&user).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "users Created"})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
