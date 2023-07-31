package controller

import (
	"api/internal/auth"
	"api/internal/database"
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/security"

	"github.com/gofiber/fiber/v2"
)

// Login is responsible for auth of the users
func Login(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	tempUser, err := repository.SearchByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := security.ComparePasswordWithHash(tempUser.Password, user.Password); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := auth.GenerateToken(tempUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).SendString(token)
}
