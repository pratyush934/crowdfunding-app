package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pratyush934/crowdfunding-app/user-service/models"
	"strings"
	"time"
)

var privateKey = []byte("57bdcc478aec3c27e91838f1247bc6244b12e4b63bb2a3cc62bb2b2fa15683e0")

func GenerateJWT(user models.User) (string, error) {

	tokenTTL := 1800

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.RoleId,
		"iat":  time.Now().Unix(),
		"eat":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateToken(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		fmt.Println("Error in ValidateToken:", err)
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func validateAdminRole(ctx *gin.Context) error {
	token, err := getToken(ctx)

	if err != nil {
		fmt.Println("Error exist in the validateAdminRole method in jwt.go")
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := claims["role"].(float64)

	if ok && token.Valid && userRole == 1 {
		return nil
	}

	return errors.New("there is an error while validating admin")

}

func validateUserRole(ctx *gin.Context) error {

	token, err := getToken(ctx)

	if err != nil {
		fmt.Println("Error exist while validatingUserRole in jwt.go -> 1")
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	userRole := claims["role"].(float64)

	if ok && token.Valid && userRole == 1 || userRole == 2 {
		return nil
	}

	return errors.New("error exist in validatingUserRole")
}

func currentUser(ctx *gin.Context) (models.User, error) {
	token, err := getToken(ctx)

	if err != nil {
		fmt.Println("There is an error while getting currentUser in jwt.go")
		return models.User{}, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	userId := claims["id"].(string)

	currentUser, err := models.GetUserById(userId)

	if err != nil {
		return models.User{}, err
	}
	return currentUser, nil
}

func getToken(ctx *gin.Context) (*jwt.Token, error) {
	token, err := getTokenFromHeader(ctx)

	//fmt.Println("Second time", token)

	if err != nil {
		fmt.Println("Error exist in the getToken method in jwt.go")
		return nil, err
	}

	//verifying
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}

func getTokenFromHeader(ctx *gin.Context) (string, error) {
	bearerToken := ctx.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	//fmt.Println(splitToken[1], "------->This is the first time")

	if len(splitToken) != 2 {
		return "", nil
	}
	return splitToken[1], nil
}
