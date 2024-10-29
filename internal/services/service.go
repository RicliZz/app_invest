package services

import (
	"github.com/RicliZz/app_invest/internal/models/authModel"
)

type AuthServiceInterface interface {
	SignUp(payload authModel.RequestSignUpPayload) error
	SignIn(payload authModel.RequestSignInPayload) (*authModel.ResponseSignInPayload, error)
}
