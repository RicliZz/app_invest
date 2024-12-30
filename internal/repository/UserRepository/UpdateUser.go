package UserRepository

import (
	"github.com/RicliZz/app_invest/internal/models/requests"
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"gorm.io/gorm"
)

func (rep *UserRepositoryImpl) SaveUpdates(user userModel.User) error {
	err := rep.db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (rep *UserRepositoryImpl) UpdateUserAndDetailsTransaction(userID int64, updates requests.PatchProfileRequest) error {
	return rep.db.Transaction(func(tx *gorm.DB) error {
		var user userModel.User
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}
		if err := tx.Model(&user).Updates(updates.UpdateUserRequest).Error; err != nil {
			return err
		}
		var userDetails userModel.UserDetails
		if err := tx.First(&userDetails, "user_id = ?", userID).Error; err != nil {
			return err
		}
		if err := tx.Model(&userDetails).Updates(updates.UpdateDetailsRequest).Error; err != nil {
			return err
		}
		return nil
	})
}
