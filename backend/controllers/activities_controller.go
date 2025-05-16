package controllers

import (
	"backend/services"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHotelByID(ctx *gin.Context) {
	id := ctx.Param("id")
	activityID, err := strconv.Atoi(id)
	if err != nil {
		ctx.String(http.StatusBadRequest, "ID invalido")
		return
	}
	hotel := services.GetActivityByID(activityID)
	ctx.JSON(http.StatusOK, hotel)
}
