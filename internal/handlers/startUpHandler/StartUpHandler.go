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
			suIDRouter.GET("/", func(c *gin.Context) {

			})

			suIDRouter.PATCH("/update", func(c *gin.Context) {

			})

			suIDRouter.DELETE("/delete", func(c *gin.Context) {

			})

		}

		startUpRouter.GET("/search-all", func(c *gin.Context) {})

		startUpRouter.POST("/create", h.service.CreateNewStartUp)
	}
}
