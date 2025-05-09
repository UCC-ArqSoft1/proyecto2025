package main

import (
	"backend/controllers"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func main() {

	services.Login("emiliano", "1234")

	router := gin.New()
	router.GET("/hotels/:id", controllers.GetHotelByID)
	router.Run()
}
