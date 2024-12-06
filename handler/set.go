package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetHandler(ctx *gin.Context) {

	request := SetContractValueRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
}
