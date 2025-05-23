package main

import (
	"backend/controllers"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.POST("/users/login", utils.CORS, controllers.Login)
	router.GET("/activities/:id", utils.CORS, controllers.GetHotelByID)
	router.Run()
}
