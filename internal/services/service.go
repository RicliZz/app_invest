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
	CreateStartUp(ctx *gin.Context)
	DeleteStartUp(ctx *gin.Context)
	UpdateStartUp(ctx *gin.Context)
	GetStartUp(ctx *gin.Context)
	GetAllStartUps(ctx *gin.Context)
}
