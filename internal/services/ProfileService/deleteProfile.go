package profileService

import (
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary deleteProfile
// @Description Удалить профиль
// @Tags Profile
// @Accept  json
// @Produce  json
// @Param id path int true "ID пользователя"
// @Success 200 {string} string ""
// @Failure 400 {string} string ""
// @Security BearerAuth
// @Router /profile/delete [delete]
func (s *ProfileService) DeleteProfile(c *gin.Context) {
	userID, _ := Utils.GetUserFromContext(c)
	err := s.repoUser.DeleteUser(userID)
	if err != nil {
		c.JSON(500, "Not deleted")
		return
	}
	c.JSON(200, "Deleted")
}
