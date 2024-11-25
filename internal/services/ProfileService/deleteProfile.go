package profileService

import (
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
)

func (s *ProfileService) DeleteProfile(c *gin.Context) {
	userID := Utils.GetUserIDFromContext(c)
	err := s.repoUser.DeleteUser(userID)
	if err != nil {
		c.JSON(500, "Not deleted")
		return
	}
	c.JSON(200, "Deleted")
}
