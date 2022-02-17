package routes

import (
	"fiber/app/controllers/postControllers"
	"fiber/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(api *fiber.App) {
	auth := api.Group("/api/posts")
	auth.Get("/", middlewares.Protected(), postControllers.Index)
	auth.Get("/show/:id", middlewares.Protected(), postControllers.ShowPost)
	auth.Post("/create", middlewares.Protected(), postControllers.StorePost)
	auth.Put("/update/:id", middlewares.Protected(), postControllers.UpdatePost)
	auth.Delete("/delete/:id", middlewares.Protected(), postControllers.DeletePost)
}
