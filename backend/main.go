package main

import (
	"backend/controllers"
	"backend/services"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	services.Login("emiliano", "1234")

	router := gin.New()
	router.POST("/users/login", utils.CORS, controllers.Login)
	router.GET("/activities/:id", utils.CORS, controllers.GetHotelByID)
	router.Run()
}
