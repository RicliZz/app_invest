package authService

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s *AuthService) VerifyToken(c *gin.Context) {
	emailToken := c.Param("emailToken")
	if emailToken == "" {
		log.Println("token required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
		return
	}

	email, err := s.redisClient.Get(c.Request.Context(), emailToken).Result()
	if err != nil {
		log.Println(err)
	}

	user, err := s.repoUser.GetUserByEmail(email)
	if err != nil {
		log.Println(err, "Not find user")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user.EmailConfirm = true
	err = s.repoUser.SaveUpdates(*user)
	if err != nil {
		log.Println(err, "Not save user")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, "Успешно!")
}
