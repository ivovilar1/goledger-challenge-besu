package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
