package middleware

import (
	"fmt"
	"net/http"
	"time"
	"vulcan/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// enforce signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return "", fmt.Errorf("token parse failed: %w", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username missing or invalid")
	}

	return username, nil
}

func LoginHandler(context *gin.Context) {
	var req models.LoginRequest

	if err := context.ShouldBindBodyWithJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	if req.Username != "admin" || req.Password != "admin" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})

		return
	}

	token, err := createToken(req.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
