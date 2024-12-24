package AdminRepository

import "github.com/RicliZz/app_invest/internal/models/userModel"

func (rep *Admin) GetUsersWithStartups() ([]userModel.User, error) {
	var users []userModel.User
	res := rep.db.Joins("JOIN startup ON startup.userid = users.id").Group("users.id").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}
