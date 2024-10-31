package authRepository

import (
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"gorm.io/gorm"
	"log"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db}
}

func (repAuth *AuthRepositoryImpl) Create(payload authModel.RequestSignUpPayload) error {
	newUser := userModel.User{
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		MiddleName: payload.MiddleName,
		Email:      payload.Email,
		Password:   payload.Password,
	}
	result := repAuth.db.Create(&newUser)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	log.Printf("Successfully created user with ID = %d", newUser.ID)
	return nil
}
