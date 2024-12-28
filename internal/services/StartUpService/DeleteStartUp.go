package StartUpService

import (
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @BasePath /api/v1

// @Summary deleteStartup
// @Description Удаление стартапа
// @Tags StartUp
// @Accept  json
// @Produce  json
// @Param id path int true "ID стартапа"
// @Security BearerAuth
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /startup/{id}/delete [delete]
func (s *StartUpService) DeleteStartUp(c *gin.Context) {
	startUpId, _ := Utils.GetUserFromContext(c)

	_, err := s.repoStartUp.GetStartUpById(startUpId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		log.Println(err)
		return
	}

	err = s.repoStartUp.Delete(startUpId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": startUpId})
}
