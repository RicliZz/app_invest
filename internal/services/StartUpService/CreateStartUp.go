package StartUpService

import (
	"fmt"
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"github.com/RicliZz/app_invest/internal/models/requests"
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StartUpService struct {
	repoStartUp     repository.StartUpRepository
	repoUser        repository.UserRepository
	repoUserDetails repository.UserDetailsRepository
}

func NewStartUpService(repoStartUp repository.StartUpRepository, repoUser repository.UserRepository, repoUserDetails repository.UserDetailsRepository) *StartUpService {
	return &StartUpService{
		repoStartUp:     repoStartUp,
		repoUser:        repoUser,
		repoUserDetails: repoUserDetails,
	}
}

// @BasePath /api/v1

// @Summary create startup
// @Description Создание стартапа
// @Tags StartUp
// @Accept  json
// @Produce  json
// @Param body body requests.StartupRequest  true  "Данные для создания стартапа"
// @Security BearerAuth
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /startup/create [post]
func (s *StartUpService) CreateStartUp(c *gin.Context) {
	var payload requests.StartupRequest
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.Title == nil || payload.Topic == nil || payload.FundingGoal == nil || payload.OfferedPercent == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload, no mandatory fields"})
	}

	if s.repoStartUp.IsTitleExists(*payload.Title) {
		res := fmt.Sprintf("title with name %s already exists", payload.Title)
		c.JSON(http.StatusBadRequest, gin.H{"error": res})
		return
	} else {
		userId, _ := Utils.GetUserFromContext(c)

		startup := startUpModel.StartUp{
			UserID:            int(userId),
			Title:             *payload.Title,
			Topic:             *payload.Topic,
			Idea:              *payload.Idea,
			Strategy:          *payload.Strategy,
			HistoryOfCreation: *payload.HistoryOfCreation,
			Status:            *payload.Status,
			Stage:             *payload.Stage,
			FundingGoal:       *payload.FundingGoal,
			OfferedPercent:    *payload.OfferedPercent}

		err := s.repoStartUp.Create(startup)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "startup created successfully")
	}
}
