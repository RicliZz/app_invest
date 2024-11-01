package services

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	userProfileModel "github.com/RicliZz/app_invest/internal/models/userProfile"
)

type AuthServiceInterface interface {
	SignUp(payload authModel.RequestSignUpPayload) error
	SignIn(payload authModel.RequestSignInPayload) (*authModel.ResponseSignInPayload, error)
}

type ProfileServiceInterface interface {
	GetProfile(userId int64) (*userProfileModel.UserProfile, error)
}

type StartUpServiceInterface interface {
	CreateNewStartUp(payload startUpModel.StartUp) error
}
