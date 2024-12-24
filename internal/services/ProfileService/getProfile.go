package profileService

import (
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"log"
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
// @Description Посмотреть СВОЙ профиль
// @Tags Profile
// @Accept  json
// @Produce  json
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Security BearerAuth
// @Router /profile/ [get]
func (s *ProfileService) GetMyProfile(c *gin.Context) {

	userId, _ := Utils.GetUserFromContext(c)

	user, err := s.repoUser.GetUserById(userId)
	if err != nil {
		log.Println(err)
		return
	}
	userDetails, err := s.repoDetails.GetUserDetails(userId)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(200, &userModel.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		MiddleName:  user.MiddleName,
		Email:       user.Email,
		UserDetails: userDetails,
	})
}
