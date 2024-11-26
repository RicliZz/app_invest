package StartUpRepository

import startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"

func (rep *StartUpRepositoryImpl) GetAllStartUps( /*фильтры и пагинация*/ ) ([]startUpModel.StartUp, error) {
	startup := make([]startUpModel.StartUp, 0)
	result := rep.db.Find(&startup)
	if result.Error != nil {
		return nil, result.Error
	}
	return startup, nil
}
