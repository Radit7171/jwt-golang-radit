package main

import (
    "booking-bengkel-smkn1depok/database"
    "booking-bengkel-smkn1depok/models"
    "booking-bengkel-smkn1depok/routes"
    "github.com/gofiber/fiber/v2"
    "os"
)

func main() {
    app := fiber.New()

    // koneksi database
    database.ConnectDB()

    // migrasi tabel
    database.DB.AutoMigrate(&models.Booking{}, &models.Servis{}, &models.Teknisi{}, &models.User{})

    // daftar routes
    routes.SetupRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    app.Listen(":" + port)
}
