package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fahmilukis/expense-tracker/pkg/middlewares"
	"github.com/fahmilukis/expense-tracker/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	config := fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}

	url := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	app := fiber.New(config)

	// middleware
	middlewares.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	if err := app.Listen(url); err != nil {
		log.Printf("server is not running! Reason: %v", err)
	}
}
