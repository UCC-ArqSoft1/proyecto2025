package controllers

import (
	"backend/services"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHotelByID(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	hotelIDString := ctx.Param("id")
	hotelIDInt, err := strconv.Atoi(hotelIDString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "ID invalido")
		return
	}
	hotel := services.GetHotelByID(hotelIDInt)
	ctx.JSON(http.StatusOK, hotel)
}
