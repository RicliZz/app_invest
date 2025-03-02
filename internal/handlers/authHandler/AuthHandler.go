package authHandler

import (
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthServiceInterface
}

func NewAuthHandler(service services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/sign-up", h.service.SignUp)

		authRouter.POST("/sign-in", h.service.SignIn)

		authRouter.GET("/verify/:emailToken", h.service.VerifyToken)

		authRouter.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}
}
