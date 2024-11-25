package repository

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"github.com/google/uuid"
)

type AuthRepository interface {
	Create(payload authModel.RequestSignUpPayload, emailToken uuid.UUID) (int, error)
}

type UserRepository interface {
	IsEmailExists(email string) bool
	GetUserById(id int64) (*userModel.User, error)
	GetUserByEmail(email string) (*userModel.User, error)
	GetUserByEmailToken(token uuid.UUID) (*userModel.User, error)
	SaveUpdates(user userModel.User) error
	DeleteUser(userID int64) error
}

type UserDetailsRepository interface {
	CreateDetailsByUserId(userId int64) error
	GetUserDetails(userId int64) (*userModel.UserDetails, error)
}

type StartUpRepository interface {
	Create(payload startUpModel.StartUp) error
	IsTitleExists(title string) bool
}
