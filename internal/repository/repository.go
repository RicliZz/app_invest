package repository

import (
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/models/userModel"
)

type AuthRepository interface {
	Create(payload authModel.RequestSignUpPayload) (int, error)
}

type UserRepository interface {
	IsEmailExists(email string) bool
	GetUserById(id int64) (*userModel.User, error)
	GetUserByEmail(email string) (*userModel.User, error)
}

type UserDetailsRepository interface {
	CreateDetailsByUserId(userId int64) error
	GetUserDetails(userId int64) (*userModel.UserDetails, error)
}
