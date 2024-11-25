package authService

import (
	"errors"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"net/http"
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

func (s *AuthService) SignUp(c *gin.Context) {
	var payload authModel.RequestSignUpPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Проверка валидатором
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(payload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Проверка на уже существующий Email
	if s.repoUser.IsEmailExists(payload.Email) {
		log.Println(errors.New("Email already exists"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hash, err := Utils.HashPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload.Password = hash
	emailToken := uuid.New()

	//Создание нового пользователя
	userID, err := s.repoAuth.Create(payload, emailToken)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Utils.SendEmail(payload, emailToken)
	//Создание деталей для пользователя
	if err = s.repoDetails.CreateDetailsByUserId(int64(userID)); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, "Пожалуйста, проверьте свою почту для завершения регистрации.")
}
