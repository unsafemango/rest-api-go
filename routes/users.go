package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"unsafemango.com/rest-api-go/models"
	"unsafemango.com/rest-api-go/utils"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data.",
		})
		return
	}

	userId, err := user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user. Email already exists.",
		})
		return
	} else {
		fmt.Println("User ID: ", userId)
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data.",
		})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user.",
		})
		return
	}

	// generate token for user
	token, err := utils.GenrateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login Successful.",
		"token":   token,
	})
}
