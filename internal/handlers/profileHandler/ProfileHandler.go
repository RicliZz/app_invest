package profileHandler

import (
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProfileHandler struct {
	service services.ProfileServiceInterface
}

func NewProfileHandler(service services.ProfileServiceInterface) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) RegisterRoutes(router *gin.Engine) {
	profileRouter := router.Group("/profile")
	{
		profileRouter.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			id64 := int64(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			response, err := h.service.GetProfile(id64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, response)
		})
	}
}
