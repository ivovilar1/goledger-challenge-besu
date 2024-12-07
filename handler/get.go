package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	value, err := CallContract()
	if err != nil {
		logger.Errf("error to retrieve value of contract: %v", err)
		SendError(ctx, http.StatusBadRequest, "error to retrieve value of contract")
		return
	}

	SendSucess(ctx, "operation-set", value)
}
