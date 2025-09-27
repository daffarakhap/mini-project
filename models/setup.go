package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file not found, using system env")
	}

	// Ambil dari env
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		host := os.Getenv("PGHOST")
		user := os.Getenv("PGUSER")
		password := os.Getenv("PGPASSWORD")
		dbname := os.Getenv("PGDATABASE")
		port := os.Getenv("PGPORT")
		sslmode := os.Getenv("PGSSLMODE")
		if sslmode == "" {
			sslmode = "disable" // default kalau tidak ada
		}

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
			host, user, password, dbname, port, sslmode,
		)
	}

	// Connect ke PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Database connected successfully!")
}



