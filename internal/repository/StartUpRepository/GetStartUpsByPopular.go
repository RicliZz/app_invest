package StartUpRepository

import startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"

func (rep *StartUpRepositoryImpl) GetStartUpsByPopularity() ([]startUpModel.StartUp, error) {
	startUps := make([]startUpModel.StartUp, 0)
	rep.db.Order("popularity asc").Limit(5).Find(startUps)
	if rep.db.Error != nil {
		return startUps, rep.db.Error
	}
	return startUps, nil
}
