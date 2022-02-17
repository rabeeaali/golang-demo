package routes

import (
	"fiber/app/middlewares"
	"github.com/gofiber/fiber/v2"
	"os"
)

func SetupRoutes() {
	app := fiber.New(fiber.Config{
		ServerHeader: os.Getenv("Server_Header"),
		AppName:      os.Getenv("APP_NAME"),
	})

	// Middlewares.
	middlewares.FiberMiddleware(app)

	// Routes.
	AuthRoutes(app)
	UserRoutes(app)
	PostRoutes(app)

	// route not found
	NotFoundRoute(app)

	// get port of server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	app.Listen(":" + port)
}
