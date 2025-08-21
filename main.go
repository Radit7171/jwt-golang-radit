package main

import (
    "booking-bengkel-smkn1depok/database"
    "booking-bengkel-smkn1depok/models"
    "booking-bengkel-smkn1depok/routes"
    "github.com/gofiber/fiber/v2"
    "os"
    "fmt"
)

func main() {
    app := fiber.New()

    // koneksi database (ambil dari env)
    database.ConnectDB()

    // migrasi tabel
    database.DB.AutoMigrate(&models.Booking{}, &models.Servis{}, &models.Teknisi{}, &models.User{})

    // daftar routes
    routes.SetupRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    fmt.Println("Server running on port", port)
    app.Listen("0.0.0.0:" + port) // <- penting
}
