package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_GO/utlis"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no token found"})
		return
	}

	userID, ok := utlis.VerifyToken(token)

	if ok != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "NOT_AUTHORIZED"})
		return
	}
	context.Set("user_id", userID)
	context.Next()
}
