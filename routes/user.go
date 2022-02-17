package routes

import (
	"fiber/app/controllers/userControllers"
	"fiber/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api *fiber.App) {
	auth := api.Group("/api/user")
	auth.Get("/", middlewares.Protected(), userControllers.UserInfo)
	//auth.Get("/logout", middlewares.Protected(), userControllers.Logout)
	//auth.Patch("/update", userControllers.up)
}
