package StartUpService

import (
	"fmt"
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type StartUpService struct {
	repoStartUp repository.StartUpRepository
}

func NewStartUpService(repoStartUp repository.StartUpRepository) *StartUpService {
	return &StartUpService{
		repoStartUp: repoStartUp,
	}
}

func (s *StartUpService) CreateNewStartUp(c *gin.Context) {
	var payload startUpModel.StartUp
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(payload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("required fields are missing")})
		return
	}

	if s.repoStartUp.IsTitleExists(payload.Title) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("title with that name already exists", payload.Title)})
		return
	} else {
		if err := s.repoStartUp.Create(payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "startUp created successfully"})
			return
		}

		//добавить добавление поля userid

	}

}
