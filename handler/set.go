package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}