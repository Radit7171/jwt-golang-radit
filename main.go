package main

import (
	"booking-bengkel-smkn1depok/database"
	"booking-bengkel-smkn1depok/models"
	"booking-bengkel-smkn1depok/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// koneksi database
	database.ConnectDB()

	// migrasi tabel
	database.DB.AutoMigrate(&models.Booking{}, &models.Servis{}, &models.Teknisi{}, &models.User{})

	// init fiber
	app := fiber.New()

	// daftar routes
	routes.SetupRoutes(app)

	// ambil PORT dari Render / default 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server running on port %s 🚀\n", port)
	app.Listen(":" + port)
}
