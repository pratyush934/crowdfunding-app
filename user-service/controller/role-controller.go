package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/crowdfunding-app/user-service/models"
	"net/http"
	"strconv"
)

func AssignRole(ctx *gin.Context) {

	var role models.Role

	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "There is an error in the AssignRole Method in role-controller.go",
		})
		ctx.Abort()
		return
	}
	err = models.CreateRole(&role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "There is error in creating role",
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Role Created Successfully",
		"role":    role,
	})
}

func GetRoles(ctx *gin.Context) {
	var roles []models.Role

	err := models.GetRoles(&roles)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "There is an error while getting the roles from GetRoles method",
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, roles)
}

func GetRole(ctx *gin.Context) {
	var role models.Role

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := models.GetRole(&role, int64(id))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "There is an error in the GetRoleMethod",
			"error":        err,
		})
		return
	}

	ctx.JSON(http.StatusOK, role)
}

func UpdateRole(ctx *gin.Context) {
	var role models.Role

	err := models.UpdateRole(&role)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "There is an error in the UpdateRole",
			"error":        err,
		})
		return
	}

	ctx.JSON(http.StatusOK, role)

}
