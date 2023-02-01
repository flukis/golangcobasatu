package controllers

import (
	"fmt"
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

type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
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

	if !utils.SimpleEmailValidation(payload.email) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("email %s in wrong format", payload.email),
		})
	}

	if !utils.SimplePasswordValidation(payload.password) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "password atleast have 1 uppercase letter, 1 special character, and min 8 character",
		})
	}

	// check if email already regisered
	_, err = db.GetAccountByEmail(payload.email)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  false,
			"message": fmt.Errorf("email %s already registered", payload.email),
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

func LoginUser(c *fiber.Ctx) error {
	signIn := &SignIn{}

	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	foundedUser, err := db.GetAccountByEmail(signIn.Email)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given email is not found",
		})
	}

	comparedPassword := utils.ComparePassword(foundedUser.Password, signIn.Password)
	if !comparedPassword {
		// Return, if password is not compare to stored in database.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "wrong user email address or password",
		})
	}

	tokens, err := utils.GenerateNewTokens(foundedUser.ID.String())
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// return success
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success login account",
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}
