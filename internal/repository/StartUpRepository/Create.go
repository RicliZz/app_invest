package StartUpRepository

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"gorm.io/gorm"
)

type StartUpRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *StartUpRepositoryImpl { return &StartUpRepositoryImpl{db} }

func (rep *StartUpRepositoryImpl) Create(payload startUpModel.StartUp) error {
	return nil
}
