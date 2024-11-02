package services

import (
	"github.com/RicliZz/app_invest/internal/models/authModel"
	userProfileModel "github.com/RicliZz/app_invest/internal/models/userProfile"
	"github.com/gin-gonic/gin"
)

type AuthServiceInterface interface {
	SignUp(payload authModel.RequestSignUpPayload) error
	SignIn(payload authModel.RequestSignInPayload) (*authModel.ResponseSignInPayload, error)
}

type ProfileServiceInterface interface {
	GetProfile(userId int64) (*userProfileModel.UserProfile, error)
}

type StartUpServiceInterface interface {
	CreateNewStartUp(ctx *gin.Context)
}
