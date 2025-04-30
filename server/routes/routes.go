package routes

import (
	controller "Chat-Application/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) { // Test route to verify application setup
	app.Get("/", controller.Hello)
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	// app.Get("/api/user", controllers.User)
	app.Post("/logout", controller.Logout)
}
