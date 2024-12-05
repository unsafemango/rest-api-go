package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"unsafemango.com/rest-api-go/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	// check if we have a token
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	// check if its not a valid token
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not unauthorized",
		})
		return
	}

	context.Set("userId", userId)

	context.Next() // ensures the next request handler will execute correctly
}
