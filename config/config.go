// config/config.go
package config

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

// ConnectDatabase buka koneksi ke SQLite
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("booking_bengkel.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}

	DB = database
	fmt.Println("Database connected 🚀")
}