package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Panic("❌ DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("❌ Failed to connect to database: %v", err)
	}

	// ✅ Simpan koneksi
	DB = db

	// ✅ AutoMigrate semua model
	err = db.AutoMigrate(&Bioskop{})
	if err != nil {
		log.Panicf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Database connected & migrated")
}




