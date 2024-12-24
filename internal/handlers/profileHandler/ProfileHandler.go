package profileHandler

import (
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	service services.ProfileServiceInterface
}

func NewProfileHandler(service services.ProfileServiceInterface) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) RegisterRoutes(router *gin.RouterGroup) {
	profileRouter := router.Group("/profile")
	profileRouter.Use(Utils.AuthMiddleware())
	{
		profileRouter.GET("/:id", h.service.GetProfile)
		profileRouter.GET("/statistic", h.service.GetProfile)
		profileRouter.PATCH("/update", h.service.PatchProfile)
		profileRouter.DELETE("/delete", h.service.DeleteProfile)
	}
}
