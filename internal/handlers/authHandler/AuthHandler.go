package authHandler

import (
	"github.com/RicliZz/app_invest/internal/models/authModel"
	"github.com/RicliZz/app_invest/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service services.AuthServiceInterface
}

func NewAuthHandler(service services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/sign-up", func(c *gin.Context) {
			var signUpParams authModel.RequestSignUpPayload
			if err := c.ShouldBindJSON(&signUpParams); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}
			if err := h.service.SignUp(signUpParams); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
		})

		authRouter.POST("/sign-in", func(c *gin.Context) {
			var signInParams authModel.RequestSignInPayload
			if err := c.ShouldBindJSON(&signInParams); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}
			//if err := h.service.SignIn(signInParams); err != nil {
			//	c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			//	return
			//}

			response, err := h.service.SignIn(signInParams)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{"token": response.Jwt})
		})
	}
}
