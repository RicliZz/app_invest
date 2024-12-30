package profileService

import (
	"github.com/RicliZz/app_invest/internal/models/requests"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary patchProfile
// @Description Обновить профиль
// @Tags Profile
// @Accept  json
// @Produce  json
// @Param body body requests.PatchProfileRequest true  "Данные для обновления профиля"
// @Success 200 {string} string ""
// @Failure 400 {string} string ""
// @Security BearerAuth
// @Router /profile/update [patch]
func (s *ProfileService) PatchProfile(c *gin.Context) {
	userID, _ := Utils.GetUserFromContext(c)

	var updates requests.PatchProfileRequest

	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.repoUser.UpdateUserAndDetailsTransaction(userID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User successfully updated")
}
