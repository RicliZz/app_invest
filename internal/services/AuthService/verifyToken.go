package authService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (s *AuthService) VerifyToken(c *gin.Context) {
	token := c.Param("emailToken")
	if token == "" {
		log.Println("token required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
		return
	}

	uuidToken, err := uuid.Parse(token)

	user, err := s.repoUser.GetUserByEmailToken(uuidToken)
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
}
