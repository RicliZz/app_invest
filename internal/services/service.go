package services

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"github.com/RicliZz/app_invest/internal/models/authModel"
)

type AuthServiceInterface interface {
	SignUp(payload authModel.RequestSignUpPayload) error
	SignIn(payload authModel.RequestSignInPayload) (*authModel.ResponseSignInPayload, error)
}

type StartUpServiceInterface interface {
	CreateNewStartUp(payload startUpModel.StartUp) error
}
