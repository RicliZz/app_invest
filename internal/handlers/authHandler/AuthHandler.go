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
		// @Summary Auth
		// @Description Пользователь регистрируется в системе, предоставив необходимые данные.
		// @Tags auth
		// @Accept  json
		// @Produce  json
		// @Param   body  body  SignUpRequest  true  "Данные для регистрации"
		// @Success 201 {object} TokenResponse
		// @Failure 400 {object} ErrorResponse
		// @Router /auth/sign-up [post]
		authRouter.POST("/sign-up", h.service.SignUp)

		authRouter.POST("/sign-in", h.service.SignIn)

		authRouter.GET("/verify/:emailToken", h.service.VerifyToken)
	}
}
