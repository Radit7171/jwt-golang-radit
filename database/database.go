package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// ambil path db dari .env
	dbPath := os.Getenv("DATABASE_URL")
	if dbPath == "" {
		dbPath = "booking_bengkel.db" // fallback kalau gak ada
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	DB = database
}
