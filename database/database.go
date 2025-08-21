package database

import (
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	// ganti import driver sqlite dengan modernc
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	dbPath := os.Getenv("DATABASE_URL")
	if dbPath == "" {
		dbPath = "booking_bengkel.db"
	}

	// pakai pure go sqlite driver
	database, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        dbPath,
	}, &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	DB = database
}
