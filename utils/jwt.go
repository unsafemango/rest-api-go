package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkey"

func GenrateToken(email string, userId int64) (string, error) {
	// generate a new token with data attached to it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // make token valid for two hours
	})

	// return a single string to be sent to the client
	return token.SignedString([]byte(secretKey))
}
