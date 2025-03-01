package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := ValidateToken(context)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Sorry But you token is not valid",
			})
			context.Abort()
			return
		}

		issue := validateAdminRole(context)

		if issue != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "You are not admin as per the token",
			})
			context.Abort()
			return
		}

		context.Next()
	}

}

func JWTAuthCustomerRole() gin.HandlerFunc {
	return func(context *gin.Context) {

		err := ValidateToken(context)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error ": "Sorry But the token is not valid",
			})
			context.Abort()
			return
		}

		issue := validateUserRole(context)

		if issue != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "You are not a customer who is registered",
			})
			context.Abort()
			return
		}
	}
}
