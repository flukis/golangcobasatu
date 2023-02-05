package routes

import (
	"github.com/fahmilukis/expense-tracker/app/controllers"
	"github.com/fahmilukis/expense-tracker/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/expense-category", middlewares.JWTProtected(), controllers.GetExpenseCategory)
	route.Get("/expense-category/:id", middlewares.JWTProtected(), controllers.GetExpenseCategoryById)
}
