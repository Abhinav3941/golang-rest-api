package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_GO/models"
	"rest_api_GO/utlis"
)

// contain all the user realted requests handlerss

// LOgin:
func Login(context *gin.Context) {
	//validate user data

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"incorrct credentials": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"can not validate ": err.Error()})
		return
	}

	token, err := utlis.GenerateJWT(user.Email, user.User_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"could not generate token": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"user is authenticated": "success", "token": token, "user infomation": user})
}

func Signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"cant save user": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user saved successfully"})
}
