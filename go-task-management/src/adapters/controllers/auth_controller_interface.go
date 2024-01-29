package controllers

import "github.com/gin-gonic/gin"

type IAuthController interface {
	Register(ctx *gin.Context)
	// Login(ctx *gin.Context)
	// RefreshToken(ctx *gin.Context)
	// Logout(ctx *gin.Context)
}
