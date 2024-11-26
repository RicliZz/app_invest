package UserRepository

import (
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"gorm.io/gorm"
	"log"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (rep *UserRepositoryImpl) IsEmailExists(email string) bool {
	var findUser userModel.User
	rep.db.Find(&findUser, "email = ?", email)
	if findUser.ID == 0 {
		return false
	}
	return true
}

func (rep *UserRepositoryImpl) GetUserById(id int64) (*userModel.User, error) {
	var user userModel.User

	result := rep.db.First(&user, "id = ?", id)

	if result.Error != nil {
		log.Println("Error with find user")
		return nil, result.Error
	}

	return &user, nil
}

func (rep *UserRepositoryImpl) GetUserByEmail(email string) (*userModel.User, error) {
	var user userModel.User

	result := rep.db.First(&user, "email = ?", email)

	if result.Error != nil {
		log.Println("Error with find user")
		return nil, result.Error
	}

	return &user, nil
}
