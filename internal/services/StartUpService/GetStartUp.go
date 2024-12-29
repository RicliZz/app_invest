package StartUpService

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary getStartup
// @Description Просмотреть стартап
// @Tags StartUp
// @Accept  json
// @Produce  json
// @Param id path int true "ID стартапа"
// @Security BearerAuth
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /startup/{id}/ [get]
func (s *StartUpService) GetStartUp(c *gin.Context) {
	startUpId := Utils.GetIDFromContext(c)
	userId, _ := Utils.GetUserFromContext(c)

	startUp, err := s.repoStartUp.GetStartUpById(startUpId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	if startUp.UserID == int(userId) {
		c.JSON(http.StatusOK, startUp)
	} else {
		user, _ := s.repoUser.GetUserById(int64(startUp.UserID))
		userDetails, _ := s.repoUserDetails.GetUserDetails(int64(startUp.UserID))
		startUp.FounderFullName = user.LastName + " " + user.FirstName + " " + user.MiddleName
		startUp.FounderEmail = user.Email
		startUp.FounderSocials = startUpModel.FounderSocials(userDetails.Socials)
		c.JSON(http.StatusOK, startUp)
	}

}
