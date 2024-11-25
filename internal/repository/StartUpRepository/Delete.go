package StartUpRepository

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
)

func (rep *StartUpRepositoryImpl) Delete(startUpId int64) error {
	var startup startUpModel.StartUp
	result := rep.db.Delete(&startup, startUpId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
