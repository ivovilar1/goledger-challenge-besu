package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivovilar1/goledger-challenge-besu/schemas"
)

func SyncHandler(ctx *gin.Context) {
	value, err := CallContract()
	if err != nil {
		logger.Errf("error to retrieve value: %v", err)
		SendError(ctx, http.StatusBadRequest, "error to retrieve value")
		return
	}
	valueStr := fmt.Sprintf("%v", value)
	contract := schemas.Contract{
		Value: valueStr,
	}

	if err := db.Create(&contract).Error; err != nil {
		logger.Errf("error to sync value: %v", err.Error())
		return
	}
	SendSucess(ctx, "operation-sync", contract)
}
