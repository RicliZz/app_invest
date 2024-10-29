package UserRepository

import (
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (rep *UserRepositoryImpl) FindUserByEmail(email string) bool {
	var findUser userModel.User
	rep.db.Find(&findUser, "email = ?", email)
	if findUser.ID == 0 {
		return false
	}
	return true
}
