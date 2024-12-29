package profileService

import (
	"github.com/RicliZz/app_invest/internal/models/responses"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"log"
)

// @BasePath /api/v1

// @Summary getStatistic
// @Description Посмотреть свой профиль
// @Tags Profile
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.UserResponse
// @Failure 400 {string} string "Need int ID"
// @Security BearerAuth
// @Router /profile/statistic [get]
func (s *ProfileService) GetStatistic(c *gin.Context) {

	userId, role := Utils.GetUserFromContext(c)

	user, err := s.repoUser.GetUserById(userId)
	if err != nil {
		log.Println("Failed to get user by ID:", err)
		c.JSON(500, "Failed to retrieve user")
		return
	}

	userDetails, err := s.repoDetails.GetUserDetails(userId)
	if err != nil {
		log.Println("Failed to get user details:", err)
		c.JSON(500, "Failed to retrieve user details")
		return
	}

	if role == "user" {
		c.JSON(200, responses.UserResponse{
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			MiddleName: user.MiddleName,
			Email:      user.Email,
			UserDetailsResponse: responses.UserDetailsResponse{
				Balance: userDetails.Balance,
				Socials: userDetails.Socials,
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"User": user,
	})
}
