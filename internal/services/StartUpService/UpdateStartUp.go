package StartUpService

import (
	"github.com/RicliZz/app_invest/internal/models/requests"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

// @BasePath /api/v1

// @Summary update startup
// @Description Обновление стартапа
// @Tags StartUp
// @Accept  json
// @Produce  json
// @Param id path int true "ID стартапа"
// @Param body body requests.StartupRequest true  "Данные для обновления стартапа"
// @Security BearerAuth
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /startup/{id}/update [patch]
func (s *StartUpService) UpdateStartUp(c *gin.Context) {
	Id := Utils.GetIDFromContext(c)

	var payload requests.StartupRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedStartup, err := s.repoStartUp.GetStartUpById(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if (payload.Title != nil && *payload.Title == "") || (payload.FundingGoal != nil && *payload.FundingGoal == 0) ||
		(payload.OfferedPercent != nil && *payload.OfferedPercent == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload, no mandatory fields"})
		return
	}

	//копирование данных из реквеста в основную структуру без ниловских полей
	err = copier.CopyWithOption(updatedStartup, payload, copier.Option{IgnoreEmpty: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = s.repoStartUp.Update(*updatedStartup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Startup was successfully updated"})
}
