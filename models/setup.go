package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ambil URL database dari Railway (sudah include host, user, password, dbname, port)
	dsn := os.Getenv("postgresql://postgres:inKAsaiCbAimbamplOoHsPSHGmcrNbcc@postgres.railway.internal:5432/railway")
	if dsn == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	// Connect ke database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	fmt.Println("✅ Connected to database")

	// Auto migrate models
	err = database.AutoMigrate(&Bioskop{})
	if err != nil {
		panic("Migration failed: " + err.Error())
	}

	fmt.Println("✅ Migration success")

	DB = database
}

