package StartUpService

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Получить ТОП-5 стартапов
// @Description Топ-5 стартапов
// @Tags StartUp
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /startup/most-popular [GET]
func (s *StartUpService) GetMostPopularStartUps(c *gin.Context) {
	startUps, err := s.repoStartUp.GetStartUpsByPopularity()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	c.JSON(http.StatusOK, startUps)
}
