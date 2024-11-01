package StartUpService

import (
	startUpModel "github.com/RicliZz/app_invest/internal/models/StartUp"
	"github.com/RicliZz/app_invest/internal/repository"
)

type StartUpService struct {
	repoStartUp repository.StartUpRepository
}

func NewStartUpService(repoStartUp repository.StartUpRepository) *StartUpService {
	return &StartUpService{
		repoStartUp: repoStartUp,
	}
}

func (s *StartUpService) CreateNewStartUp(payload startUpModel.StartUp) error {
	return nil
}
