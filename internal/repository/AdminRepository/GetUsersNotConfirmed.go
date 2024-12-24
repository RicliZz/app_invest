package AdminRepository

import "github.com/RicliZz/app_invest/internal/models/userModel"

func (rep *Admin) GetUsersNotConfirmed() ([]userModel.User, error) {
	var users []userModel.User
	res := rep.db.Preload("Startups").Preload("UserDetails").Find(&users, "confirmed", false)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}
