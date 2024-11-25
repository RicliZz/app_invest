package authService

import (
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func (s *AuthService) SignIn(c *gin.Context) {
	var payload authModel.RequestSignInPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(payload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	checkUser, err := s.repoUser.GetUserByEmail(payload.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if !checkUser.EmailConfirm {
		log.Println("Not confirm email address")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if !Utils.ComparePasswords(checkUser.Password, []byte(payload.Password)) {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Println(err)
		return
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := Utils.CreateJWT(secret, checkUser.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(200, gin.H{"Token": token})
}
