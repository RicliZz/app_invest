package StartUpService

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary getAllStartups
// @Description Просмотреть все стартапы
// @Tags StartUp
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /startup/search-all [get]
func (s *StartUpService) GetAllStartUps(c *gin.Context) {
	//в будущем нужно будет вытаскивать и юзать фильтры и пагинацию
	startUps, err := s.repoStartUp.GetAllStartUps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, startUps)

}
