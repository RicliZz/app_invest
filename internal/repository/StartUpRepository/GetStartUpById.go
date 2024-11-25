package StartUpRepository

import (
	"fmt"
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
)

func (rep *StartUpRepositoryImpl) GetStartUpById(startUpId int64) (*startUpModel.StartUp, error) {
	var startup startUpModel.StartUp
	if err := rep.db.First(&startup, startUpId).Error; err != nil {
		return nil, fmt.Errorf("startup not found")
	}
	return &startup, nil
}
