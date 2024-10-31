package repository

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/models/userModel"
)

type AuthRepository interface {
	Create(payload authModel.RequestSignUpPayload) error
}

type UserRepository interface {
	IsEmailExists(email string) bool
	GetUserById(id int64) (*userModel.User, error)
	GetUserByEmail(email string) (*userModel.User, error)
}

type StartUpRepository interface {
	Create(payload startUpModel.StartUp) error
}
