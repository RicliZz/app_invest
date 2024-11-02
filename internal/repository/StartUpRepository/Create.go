package StartUpRepository

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"gorm.io/gorm"
	"log"
)

type StartUpRepositoryImpl struct {
	db *gorm.DB
}

func NewStartUpRepository(db *gorm.DB) *StartUpRepositoryImpl { return &StartUpRepositoryImpl{db} }

func (rep *StartUpRepositoryImpl) Create(payload startUpModel.StartUp) error {
	result := rep.db.Create(&payload)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	log.Printf("Successfully created startup with ID = %d\n", payload.ID)
	return nil
}
