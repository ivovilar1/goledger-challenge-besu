package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SyncHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
