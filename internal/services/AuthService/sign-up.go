package authService

import (
	"errors"
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
)

type AuthService struct {
	repoAuth    repository.AuthRepository
	repoUser    repository.UserRepository
	repoDetails repository.UserDetailsRepository
	redisClient *redis.Client
}

func NewAuthService(repoAuth repository.AuthRepository, repoUser repository.UserRepository,
	repoDetails repository.UserDetailsRepository, redisClient *redis.Client) *AuthService {
	return &AuthService{
		repoAuth:    repoAuth,
		repoUser:    repoUser,
		repoDetails: repoDetails,
		redisClient: redisClient,
	}
}

// @BasePath /api/v1

// @Summary sign-up
// @Description Регистрация в системе
// @Tags Authorization
// @Accept  json
// @Produce  json
// @Param body body authModel.RequestSignUpPayload  true  "Данные для регистрации"
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /auth/sign-up [post]
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

	err = s.redisClient.Set(c.Request.Context(), emailToken.String(), payload.Email, time.Hour*24).Err()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Создание нового пользователя
	userID, err := s.repoAuth.Create(payload)
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
