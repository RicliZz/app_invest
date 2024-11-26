package StartUpService

import (
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s *StartUpService) DeleteStartUp(c *gin.Context) {
	startUpId := Utils.GetUserIDFromContext(c)

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
