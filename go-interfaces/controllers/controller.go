package controllers

import (
	"github.com/gin-gonic/gin"
	"go-interfaces/clients"
	"go-interfaces/services"
	"net/http"
)

// controller
func GetActivityByID(ctx *gin.Context) {
	activity, _ := services.GetActivityByID(1, clients.SQLite{})
	ctx.JSON(http.StatusOK, activity)
}
