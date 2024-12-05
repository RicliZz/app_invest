package infoHandler

import (
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/gin-gonic/gin"
)

type InfoHandler struct {
	service services.InfoServiceInterface
}

func NewInfoHandler(service services.InfoServiceInterface) *InfoHandler {
	return &InfoHandler{service: service}
}

func (h *InfoHandler) RegisterRoutes(router *gin.RouterGroup) {
	infoRouter := router.Group("/info")
	{
		infoRouter.GET("/about-us", h.service.GetInfo)

		infoRouter.GET("/contacts", h.service.GetContact)
	}
}
