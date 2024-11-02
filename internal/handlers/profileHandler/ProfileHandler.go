package profileHandler

import (
	"github.com/RicliZz/app_invest/internal/pkg/Utils"
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	service services.ProfileServiceInterface
}

func NewProfileHandler(service services.ProfileServiceInterface) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) RegisterRoutes(router *gin.Engine) {
	profileRouter := router.Group("/profile")
	profileRouter.Use(Utils.AuthMiddleware())
	{
		profileRouter.GET("/", h.service.GetProfile)
	}
}
