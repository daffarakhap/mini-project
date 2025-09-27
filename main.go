package main

import (
	"fmt"
	"os"

	"mini-project/controllers/bioskopcontroller"
	"mini-project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect ke database Railway
	models.ConnectDatabase()
	fmt.Println("✅ Database connected & migrated")

	// Setup Gin router
	router := gin.Default()

	// Routes
	router.GET("/bioskop", bioskopcontroller.Index)
	router.GET("/bioskop/:id", bioskopcontroller.Show)
	router.POST("/bioskop", bioskopcontroller.Create)
	router.PUT("/bioskop/:id", bioskopcontroller.Update)
	router.DELETE("/bioskop/:id", bioskopcontroller.Delete)

	// Port Railway biasanya pakai PORT dari environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default untuk local
	}

	// Jalankan server
	router.Run(":" + port)
}

