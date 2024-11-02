package UserRepository

import (
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"log"
)

func (rep *UserRepositoryImpl) SaveUpdates(user userModel.User) error {
	res := rep.db.Save(&user)
	if res.Error != nil {
		log.Println("Error with save user")
		return res.Error
	}
	return nil
}
