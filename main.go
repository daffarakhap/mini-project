package main

import (
	"mini-project/controllers/bioskopcontroller"
	"mini-project/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	models.ConnectDatabase()

	r := gin.Default()

	// Routes
	r.GET("/bioskops", bioskopcontroller.Index)
	r.GET("/bioskops/:id", bioskopcontroller.Show)
	r.POST("/bioskops", bioskopcontroller.Create)
	r.PUT("/bioskops/:id", bioskopcontroller.Update)
	r.DELETE("/bioskops/:id", bioskopcontroller.Delete)

	// Ambil port dari Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}


