package startUpHandler

import (
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/gin-gonic/gin"
)

type StartUpHandler struct {
	service services.StartUpServiceInterface
}

func NewStartUpHandler(service services.StartUpServiceInterface) *StartUpHandler {
	return &StartUpHandler{service: service}
}

func (h *StartUpHandler) RegisterRoutes(router *gin.Engine) {
	startUpRouter := router.Group("/startup")
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
