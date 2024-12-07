package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	value, err := CallContract()
	if err != nil {
		logger.Errf("error calling contract: %v", err)
		SendError(ctx, http.StatusBadRequest, "error calling contract")
		return
	}

	SendSucess(ctx, "operation-set", value)
}
