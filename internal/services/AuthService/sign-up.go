package authService

import (
	"errors"
	"fmt"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/go-playground/validator/v10"
	"log"
)

type AuthService struct {
	repoAuth    repository.AuthRepository
	repoUser    repository.UserRepository
	repoDetails repository.UserDetailsRepository
}

func NewAuthService(repoAuth repository.AuthRepository, repoUser repository.UserRepository,
	repoDetails repository.UserDetailsRepository) *AuthService {
	return &AuthService{
		repoAuth:    repoAuth,
		repoUser:    repoUser,
		repoDetails: repoDetails,
	}
}

func (s *AuthService) SignUp(payload authModel.RequestSignUpPayload) error {
	//Проверка валидатором
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(payload); err != nil {
		log.Println(err)
		return err
	}
	//Проверка на уже существующий Email
	if s.repoUser.IsEmailExists(payload.Email) {
		log.Println(errors.New("User already exists"))
		return errors.New(fmt.Sprintf("User with email %s already exists", payload.Email))
	}

	hash, err := Utils.HashPassword(payload.Password)
	if err != nil {
		return err
	}
	payload.Password = hash

	//Создание нового пользователя
	userID, err := s.repoAuth.Create(payload)
	if err != nil {
		log.Println(err)
		return err
	}
	//Создание деталей для пользователя
	if err = s.repoDetails.CreateDetailsByUserId(int64(userID)); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
