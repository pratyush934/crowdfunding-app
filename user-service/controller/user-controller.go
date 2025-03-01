package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/crowdfunding-app/user-service/models"
	"github.com/pratyush934/crowdfunding-app/user-service/util"
	"net/http"
)

func Register(context *gin.Context) {
	var input models.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "There is some issue related to registration process",
		})
		return
	}

	user := models.User{
		UserName: input.UserName,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   2,
	}

	save, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "there is some error while saving the user",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"user": save,
	})
}

func Login(context *gin.Context) {

	var loginInput models.Login

	if err := context.ShouldBindJSON(&loginInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "there is some error in the process of taking login Input",
		})
	}

	user, err := models.GetUserByEmail(loginInput.Email)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	issue := user.ValidatePassword(loginInput.Password)

	if issue != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Password",
		})
		return
	}

	jwt, err := util.GenerateJWT(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Not Able to generate JWT",
		})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"jwt":      jwt,
		"userName": user.UserName,
		"message":  "Login successful",
	})

}

func GetUsers(context *gin.Context) {

	var users []models.User

	err := models.GetUsers(&users)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Not able to get All the users",
		})
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusOK, users)
}

func GetUser(context *gin.Context) {

	var user models.User

	id := context.Param("id")

	err := models.GetUser(&user, id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Not able to get the User with Id",
		})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func updateUser(ctx *gin.Context) {

	var updatedUser models.User

	err := models.UpdateUser(&updatedUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Not able to Update the User",
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"updatedUser": updatedUser,
	})
}
