package authService

import (
	"errors"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/go-playground/validator/v10"
	"log"
)

type AuthService struct {
	repoAuth repository.AuthRepository
	repoUser repository.UserRepository
}

func NewAuthService(repoAuth repository.AuthRepository, repoUser repository.UserRepository) *AuthService {
	return &AuthService{
		repoAuth: repoAuth,
		repoUser: repoUser,
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
	if s.repoUser.FindUserByEmail(payload.Email) {
		log.Println(errors.New("User already exists"))
		return errors.New("Email already exists")
	}
	//Создание нового пользователя
	if err := s.repoAuth.Create(payload); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
