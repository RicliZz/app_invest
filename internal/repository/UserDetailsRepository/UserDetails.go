package UserDetailsRepository

import (
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"gorm.io/gorm"
	"log"
)

type UserDetailsRepositoryImpl struct {
	db *gorm.DB
}

func NewUserDetailsRepositoryImpl(db *gorm.DB) *UserDetailsRepositoryImpl {
	return &UserDetailsRepositoryImpl{db: db}
}

func (rep *UserDetailsRepositoryImpl) CreateDetailsByUserId(userId int64) error {
	userDetails := userModel.UserDetails{
		UserID: userId,
	}
	result := rep.db.Create(&userDetails)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (rep *UserDetailsRepositoryImpl) GetUserDetails(userId int64) (*userModel.UserDetails, error) {
	var findUserDetails userModel.UserDetails
	res := rep.db.Find(&findUserDetails, "user_id = ?", userId)
	if res.Error != nil {
		log.Println(res.Error)
		return nil, res.Error
	}
	return &findUserDetails, nil
}
