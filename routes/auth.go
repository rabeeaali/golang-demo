package routes

import (
	"fiber/app/controllers/authControllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api *fiber.App) {
	auth := api.Group("/api/auth")
	auth.Post("/register", authControllers.Register)
	auth.Post("/login", authControllers.Login)
	auth.Patch("/password/forget", authControllers.ForgetPassword)
	auth.Patch("/password/code", authControllers.CheckCode)
	auth.Patch("/password/reset", authControllers.ResetPassword)
}
