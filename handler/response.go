package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message:":  message,
		"errorCode": code,
	})
}

func SendSucess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message:": fmt.Sprintf("Operation %s successfull", operation),
		"data":     data,
	})
}
