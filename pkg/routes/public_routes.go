package routes

import (
	"github.com/fahmilukis/expense-tracker/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	routeV1 := a.Group("/api/v1")
	routeV1.Post("signup", controllers.CreateUser)
	routeV1.Post("signin", controllers.LoginUser)
}
