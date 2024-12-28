package StartUpService

import (
	"github.com/RicliZz/app_invest/internal/models/requests"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @BasePath /api/v1

// @Summary update startup
// @Description Обновление стартапа
// @Tags StartUp
// @Accept  json
// @Produce  json
// @Param id path int true "ID стартапа"
// @Param body body requests.UpdateStartupRequest true  "Данные для обновления стартапа"
// @Security BearerAuth
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /startup/{id}/update [patch]
func (s *StartUpService) UpdateStartUp(c *gin.Context) {
	Id := Utils.GetIDFromContext(c)

	var payload requests.UpdateStartupRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(Id)

	updatedStartup, err := s.repoStartUp.GetStartUpById(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if payload.Title == nil || payload.Topic == nil || payload.Status == nil ||
		payload.FundingGoal == nil || payload.OfferedPercent == nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload, no mandatory fields"})
		return
	}

	updatedStartup.Title = *payload.Title
	updatedStartup.Topic = *payload.Topic
	updatedStartup.Idea = *payload.Idea
	updatedStartup.Strategy = *payload.Strategy
	updatedStartup.HistoryOfCreation = *payload.HistoryOfCreation
	updatedStartup.Status = *payload.Status
	updatedStartup.Stage = *payload.Stage
	updatedStartup.FundingGoal = *payload.FundingGoal
	updatedStartup.OfferedPercent = *payload.OfferedPercent

	err = s.repoStartUp.Update(*updatedStartup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Startup was successfully updated"})
}
