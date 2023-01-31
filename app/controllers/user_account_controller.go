package controllers

import (
	"time"

	"github.com/fahmilukis/expense-tracker/app/models"
	"github.com/fahmilukis/expense-tracker/pkg/utils"
	"github.com/fahmilukis/expense-tracker/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserPayload struct {
	email    string
	password string
	name     string
}

func CreateUser(c *fiber.Ctx) error {
	payload := UserPayload{
		email:    c.FormValue("email"),
		password: c.FormValue("password"),
		name:     c.FormValue("name"),
	}

	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// generate hashed password
	hashedPassword, err := utils.GenerateHashedPassword(payload.password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// generate date
	today := time.Now().UTC()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// generate user
	user := models.User{
		ID:        uuid.New(),
		Email:     payload.email,
		Password:  hashedPassword,
		Name:      payload.name,
		CreatedAt: today,
		UpdatedAt: today,
	}

	err = db.CreateAccount(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// return success
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success create account",
	})
}
