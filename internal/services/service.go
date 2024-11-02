package services

import (
	"github.com/gin-gonic/gin"
)

type AuthServiceInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type ProfileServiceInterface interface {
	GetProfile(c *gin.Context)
}

type StartUpServiceInterface interface {
	CreateNewStartUp(ctx *gin.Context)
}
