package controllers

import "github.com/gin-gonic/gin"

type IUserController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	GetById(ctx *gin.Context)
	UpdateById(ctx *gin.Context)
	DeleteById(ctx *gin.Context)
}
