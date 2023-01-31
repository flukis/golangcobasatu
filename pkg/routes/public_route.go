package routes

import (
	"github.com/fahmilukis/expense-tracker/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	routeV1 := a.Group("/api/v1")

	routeV1.Get("books", controllers.GetExpenseCategory)
	routeV1.Get("book/:id", controllers.GetExpenseCategoryById)

	routeV1.Post("user", controllers.CreateUser)
}
