package repository

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
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
	SaveUpdates(user userModel.User) error
	DeleteUser(userID int64) error
	GetUsersByOpt(opt *string) (*[]userModel.User, error)
}

type UserDetailsRepository interface {
	CreateDetailsByUserId(userId int64) error
	GetUserDetails(userId int64) (*userModel.UserDetails, error)
}

type StartUpRepository interface {
	Create(payload startUpModel.StartUp) error
	IsTitleExists(title string) bool
	Delete(startUpId int64) error
	Update(payload startUpModel.StartUp) error
	GetStartUpById(startUpId int64) (*startUpModel.StartUp, error)
	GetAllStartUps() ([]startUpModel.StartUp, error) //в дальнейшем добавить на вход фильтр и пагинацию
}

type AdminRepository interface {
	GetUserById(userId int64) (*userModel.User, error)
	GetUsersWithStartups() ([]userModel.User, error)
	GetUsersNotConfirmed() ([]userModel.User, error)
}
