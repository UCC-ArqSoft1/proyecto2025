package controllers

import (
	"net/http"

	"backend/domain"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var request domain.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, token, err := services.Login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"token":   token,
	})
}
