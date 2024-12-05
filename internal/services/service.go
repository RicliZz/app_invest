package services

import (
	"github.com/gin-gonic/gin"
)

type AuthServiceInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	VerifyToken(c *gin.Context)
}

type ProfileServiceInterface interface {
	GetMyProfile(c *gin.Context)
	PatchProfile(c *gin.Context)
	DeleteProfile(c *gin.Context)
}

type StartUpServiceInterface interface {
	CreateStartUp(ctx *gin.Context)
	DeleteStartUp(ctx *gin.Context)
	UpdateStartUp(ctx *gin.Context)
	GetStartUp(ctx *gin.Context)
	GetAllStartUps(ctx *gin.Context)
}

type InfoServiceInterface interface {
	GetInfo(c *gin.Context)
	GetContact(c *gin.Context)
}

type AdminServiceInterface interface {
	AdminGetUser(c *gin.Context)
}
