package main

import (
	"backend/controllers"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func main() {

	services.Login("emiliano", "1234")

	router := gin.New()
	router.POST("/users/login", controllers.CORS, controllers.Login)
	router.GET("/activities/:id", controllers.CORS, controllers.GetHotelByID)
	router.Run()
}
