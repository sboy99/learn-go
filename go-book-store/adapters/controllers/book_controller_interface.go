package controllers

import "github.com/gin-gonic/gin"

type IBookController interface {
	CreateBook(ctx *gin.Context)
	ListBooks(ctx *gin.Context)
	GetOneBook(ctx *gin.Context)
	UpdateOneBook(ctx *gin.Context)
	DeleteOneBook(ctx *gin.Context)
}
