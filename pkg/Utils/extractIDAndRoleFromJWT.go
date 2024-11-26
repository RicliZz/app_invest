package Utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func ExtractUserFromJWT(c *gin.Context) (int64, string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return 0, "", errors.New("Authorization header is empty")
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return 0, "", err
	}
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		secret := []byte(os.Getenv("JWT_SECRET"))
		return secret, nil
	})

	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseInt(claims["userID"].(string), 10, 64)
		role := claims["role"].(string)
		if err != nil {
			return 0, "", err
		}
		return userID, role, nil
	}
	return 0, "", fmt.Errorf("Invalid token")
}
