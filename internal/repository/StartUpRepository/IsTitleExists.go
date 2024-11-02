package StartUpRepository

import (
	"errors"
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"gorm.io/gorm"
	"strings"
)

func (rep *StartUpRepositoryImpl) IsTitleExists(title string) bool {
	normalizedTitle := strings.ToLower(strings.ReplaceAll(title, " ", ""))

	// Проверка существования стартапа с таким названием
	var existingStartUp startUpModel.StartUp
	if err := rep.db.Where("LOWER(REPLACE(title, ' ', '')) = ?", normalizedTitle).First(&existingStartUp).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false // Ошибка при запросе
		}
		return false
	} else {
		return true
	}
}
