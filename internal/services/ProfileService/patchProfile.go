package profileService

import (
	"github.com/RicliZz/app_invest/internal/models/requests"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *ProfileService) PatchProfile(c *gin.Context) {
	userID, _ := Utils.GetUserFromContext(c)

	var updates requests.UpdateUserRequest

	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := s.repoUser.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if updates.FirstName != nil {
		user.FirstName = *updates.FirstName
	}
	if updates.LastName != nil {
		user.LastName = *updates.LastName
	}
	if updates.MiddleName != nil {
		user.MiddleName = *updates.MiddleName
	}

	if updates.Email != nil {
		user.Email = *updates.Email
	}

	if updates.NewPassword != nil {
		if Utils.ComparePasswords(user.Password, []byte(*updates.OldPassword)) {
			user.Password = *updates.NewPassword
		}
	}

	err = s.repoUser.SaveUpdates(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User successfully updated")
}
