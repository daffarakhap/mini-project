package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ambil variable dari Railway
	host := os.Getenv("postgres.railway.internal")
	user := os.Getenv("postgres")
	password := os.Getenv("inKAsaiCbAimbamplOoHsPSHGmcrNbcc")
	dbname := os.Getenv("railway")
	port := os.Getenv("5432")

	// Susun dsn sesuai format PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Auto migrate model
	database.AutoMigrate(&Bioskop{})

	DB = database
}
