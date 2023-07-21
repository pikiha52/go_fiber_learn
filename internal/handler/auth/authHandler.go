package authHandler

import (
	"go_fiber/config"
	"go_fiber/database"
	"go_fiber/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var inputLoginData LoginInput
	err := c.BodyParser(&inputLoginData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	db := database.DB
	var user model.User

	db.Find(&user, "username = ?", inputLoginData.Username)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	if !CheckPasswordHash(inputLoginData.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
