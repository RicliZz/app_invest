package InfoService

import (
	"github.com/RicliZz/app_invest/pkg/messages"
	"github.com/gin-gonic/gin"
)

type InfoService struct {
}

func NewInfoService() *InfoService {
	return &InfoService{}
}

// @BasePath /api/v1

// @Summary about-us
// @Description Информация о нас
// @Tags Info
// @Accept  json
// @Produce  json
// @Success 200 {string} string ""
// @Router /info/about-us [get]
func (s *InfoService) GetInfo(c *gin.Context) {
	mess := messages.AboutUsMessage()
	c.JSON(200, mess)
}

// @BasePath /api/v1

// @Summary contacts
// @Description Контакты для связи
// @Tags Info
// @Accept  json
// @Produce  json
// @Success 200 {string} string ""
// @Router /info/contacts [get]
func (s *InfoService) GetContact(c *gin.Context) {
	mess := messages.ContactMessage()
	c.JSON(200, mess)
}
