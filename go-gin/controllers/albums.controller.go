package controllers

import (
	"net/http"

	"learn/go-gin/services/albums"

	"github.com/gin-gonic/gin"
)

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums.GetAlbums())
}
