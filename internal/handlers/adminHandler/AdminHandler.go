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
	admin := router.Group("/admin")
	admin.Use(Utils.AdminOnlyMiddleware())
	{
		admin.POST("/search-one/:id", h.service.SearchOneUser)
		admin.POST("/get-all", h.service.SearchAllUsers)
		admin.POST("/as-user/:id", h.service.LoginAsAnotherUser)

	}
}
