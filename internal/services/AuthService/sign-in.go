package authService

import (
	"fmt"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	Utils2 "github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func (s *AuthService) SignIn(payload authModel.RequestSignInPayload) (*authModel.ResponseSignInPayload, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(payload); err != nil {
		log.Println(err)
		return nil, err
	}

	checkUser, err := s.repoUser.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email")
	}
	if !Utils2.ComparePasswords(checkUser.Password, []byte(payload.Password)) {
		return nil, fmt.Errorf("invalid password")
	}

	err = godotenv.Load()
	if err != nil {
		return nil, err
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := Utils2.CreateJWT(secret, checkUser.ID)
	if err != nil {
		return nil, err
	}

	return &authModel.ResponseSignInPayload{token}, nil
}
