package StartUpService

import (
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *StartUpService) GetStartUp(c *gin.Context) {
	startUpId, _ := Utils.GetUserFromContext(c)

	startUp, err := s.repoStartUp.GetStartUpById(startUpId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, startUp)
}
