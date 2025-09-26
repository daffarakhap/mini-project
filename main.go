package main

import (
	"mini-project/controllers/bioskopcontroller"
	"mini-project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/bioskop", bioskopcontroller.Index)
	router.GET("/bioskop/:id", bioskopcontroller.Show)
	router.POST("/bioskop", bioskopcontroller.Create)
	router.PUT("/bioskop/:id", bioskopcontroller.Update)
	router.DELETE("/bioskop/:id", bioskopcontroller.Delete)

	router.Run()
}
