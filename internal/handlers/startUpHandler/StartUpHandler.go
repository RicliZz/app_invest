package startUpHandler

import (
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
)

type StartUpHandler struct {
	service services.StartUpServiceInterface
}

func NewStartUpHandler(service services.StartUpServiceInterface) *StartUpHandler {
	return &StartUpHandler{service: service}
}

func (h *StartUpHandler) RegisterRoutes(router *gin.RouterGroup) {
	startUpRouter := router.Group("/startup")
	startUpRouter.Use(Utils.AuthMiddleware())
	{
		suIDRouter := startUpRouter.Group("/:id")
		{
			suIDRouter.Use(Utils.ValidateAndExtractIDMiddleware())

			suIDRouter.GET("/", h.service.GetStartUp)

			suIDRouter.PATCH("/update", h.service.UpdateStartUp)

			suIDRouter.DELETE("/delete", h.service.DeleteStartUp)

		}

		startUpRouter.GET("/search-all", h.service.GetAllStartUps)

		startUpRouter.POST("/create", h.service.CreateStartUp)

		startUpRouter.Use(Utils.AdminOnlyMiddleware()).GET("/pending", h.service.GetPendingStartUps)
	}
}
