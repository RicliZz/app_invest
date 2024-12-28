package profileService

import (
	"github.com/RicliZz/app_invest/internal/models/responses"
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type ProfileService struct {
	repoUser    repository.UserRepository
	repoDetails repository.UserDetailsRepository
}

func NewProfileService(repUser repository.UserRepository, repDetails repository.UserDetailsRepository) *ProfileService {
	return &ProfileService{repUser, repDetails}
}

// @BasePath /api/v1

// @Summary getProfile
// @Description Посмотреть чужой профиль
// @Tags Profile
// @Accept  json
// @Produce  json
// @Param id path int true "ID пользователя"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {string} string "Need int ID"
// @Security BearerAuth
// @Router /profile/{id} [get]
func (s *ProfileService) GetProfile(c *gin.Context) {

	pathUserId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Invalid path ID:", err)
		c.JSON(400, "Need int ID")
		return
	}

	userId, role := Utils.GetUserFromContext(c)

	user, err := s.repoUser.GetUserById(pathUserId)
	if err != nil {
		log.Println("Failed to get user by ID:", err)
		c.JSON(500, "Failed to retrieve user")
		return
	}

	userDetails, err := s.repoDetails.GetUserDetails(pathUserId)
	if err != nil {
		log.Println("Failed to get user details:", err)
		c.JSON(500, "Failed to retrieve user details")
		return
	}

	if role == "user" {
		log.Println(user.Role)
		if !Utils.ThisOwner(pathUserId, userId) {

			c.JSON(200, gin.H{
				"FullName": user.LastName + " " + user.FirstName + " " + user.MiddleName,
				"About":    user.AboutMe,
			})
			return
		}

		c.JSON(200, responses.UserResponse{
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			MiddleName: user.MiddleName,
			Email:      user.Email,
			UserDetailsResponse: responses.UserDetailsResponse{
				Balance: userDetails.Balance,
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"User": user,
	})
}
