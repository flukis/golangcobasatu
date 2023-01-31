package controllers

import (
	"fmt"

	"github.com/fahmilukis/expense-tracker/platform/database"
	"github.com/gofiber/fiber/v2"
)

func GetExpenseCategory(c *fiber.Ctx) error {
	// create connextion
	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// get all category
	category, err := db.GetCategories()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "categories not found",
			"count":   0,
			"data":    nil,
		})
	}

	// return success
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "categories found",
		"count":   len(category),
		"data":    category,
	})
}

func GetExpenseCategoryById(c *fiber.Ctx) error {
	// parse params
	id := c.Params("id")

	// create connextion
	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// get all category
	category, err := db.GetCategories()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("category with id %s not found", id),
			"data":    nil,
		})
	}

	// return success
	return c.JSON(fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("category with id %s found", id),
		"data":    category,
	})
}
