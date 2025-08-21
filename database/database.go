package database

import (
	"log"

	"booking-bengkel-smkn1depok/models"
	"github.com/glebarez/sqlite" // <- driver pure Go
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("booking_bengkel.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Servis{}, &models.Teknisi{}, &models.Booking{})
	DB = db
}
