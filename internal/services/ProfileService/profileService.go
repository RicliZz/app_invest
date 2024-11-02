package profileService

import (
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/RicliZz/app_invest/internal/repository"
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

func (s *ProfileService) GetProfile(c *gin.Context) {

	userId := Utils.GetUserIDFromContext(c)

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
		UserDetails: *userDetails,
	})
}
