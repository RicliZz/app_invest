package StartUpService

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *StartUpService) GetAllStartUps(c *gin.Context) {
	//в будущем нужно будет вытаскивать и юзать фильтры и пагинацию
	startUps, err := s.repoStartUp.GetAllStartUps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, startUps)

}
