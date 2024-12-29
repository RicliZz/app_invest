package StartUpRepository

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"log"
)

func (rep *StartUpRepositoryImpl) Update(payload startUpModel.StartUp) error {
	res := rep.db.Save(&payload)
	if res.Error != nil {
		log.Println("Error with updating startup")
		return res.Error
	}
	return nil
}
