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

func (rep *StartUpRepositoryImpl) GetStartUpsWithPending() ([]startUpModel.StartUp, error) {
	var startups []startUpModel.StartUp
	if err := rep.db.Where("pending = ?", true).Find(&startups).Error; err != nil {
		return nil, fmt.Errorf("startup not found")
	}
	return startups, nil
}
