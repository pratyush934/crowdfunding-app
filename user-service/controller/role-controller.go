package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/crowdfunding-app/user-service/models"
	"net/http"
)

func AssignRole(ctx *gin.Context) {

	var role models.Role

	err := ctx.ShouldBindJSON(&role)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
	}
}
