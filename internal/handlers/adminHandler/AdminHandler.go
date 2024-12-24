package adminHandler

import (
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service services.AdminServiceInterface
}

func NewAdminHandler(service services.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/admin")
	auth.Use(Utils.AdminOnlyMiddleware())
	{
		auth.POST("/search-one/:id", h.service.SearchOneUser)
		auth.POST("/get-all", h.service.SearchAllUsers)
		auth.POST("/as-user/:id", h.service.LoginAsAnotherUser)

	}
}
