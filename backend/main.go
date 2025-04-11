package main

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/hotels/:id", controllers.GetHotelByID)
	router.Run()
}
