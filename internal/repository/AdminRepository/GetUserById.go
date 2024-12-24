package AdminRepository

import (
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"gorm.io/gorm"
)

type Admin struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *Admin {
	return &Admin{db: db}
}

func (rep *Admin) GetUserById(userId int64) (*userModel.User, error) {
	var user userModel.User
	err := rep.db.Preload("Startups").Preload("UserDetails").First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
