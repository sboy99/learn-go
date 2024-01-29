package controllers

import "github.com/gin-gonic/gin"

type ITaskController interface {
	Create(ctx *gin.Context)
}
