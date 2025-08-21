package routes

import (
	"booking-bengkel-smkn1depok/handlers"
	"booking-bengkel-smkn1depok/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// auth
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// teknisi (protected)
	teknisi := app.Group("/teknisi", middlewares.Protected())
	teknisi.Post("/", handlers.CreateTeknisi)
	teknisi.Get("/", handlers.GetTeknisi)
	teknisi.Get("/:id", handlers.GetTeknisiByID)
	teknisi.Put("/:id", handlers.UpdateTeknisi)
	teknisi.Delete("/:id", handlers.DeleteTeknisi)
}
