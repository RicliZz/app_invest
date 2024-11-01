package profileService

import (
	userProfileModel "github.com/RicliZz/app_invest/internal/models/userProfile"
	"github.com/RicliZz/app_invest/internal/repository"
	"log"
)

type ProfileService struct {
	repoUser    repository.UserRepository
	repoDetails repository.UserDetailsRepository
}

func NewProfileService(repUser repository.UserRepository, repDetails repository.UserDetailsRepository) *ProfileService {
	return &ProfileService{repUser, repDetails}
}

func (s *ProfileService) GetProfile(userId int64) (*userProfileModel.UserProfile, error) {
	user, err := s.repoUser.GetUserById(userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	userDetails, err := s.repoDetails.GetUserDetails(userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &userProfileModel.UserProfile{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Patronymic: user.Patronymic,
		Email:      user.Email,
		Balance:    userDetails.Balance,
	}, nil
}
