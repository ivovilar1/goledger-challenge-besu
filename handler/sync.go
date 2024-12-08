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

	var contract schemas.Contract

	// check if exists a contract in db. If exists just update
	if err := db.First(&contract).Error; err != nil {
		if err.Error() == "record not found" {
			contract = schemas.Contract{Value: valueStr}
			if err := db.Create(&contract).Error; err != nil {
				logger.Errf("error to sync value: %v", err.Error())
				SendError(ctx, http.StatusInternalServerError, "error to sync value")
				return
			}
		} else {
			logger.Errf("error to find existing value: %v", err.Error())
			SendError(ctx, http.StatusInternalServerError, "error to find existing value")
			return
		}
	} else {
		contract.Value = valueStr
		if err := db.Save(&contract).Error; err != nil {
			logger.Errf("error to update value: %v", err.Error())
			SendError(ctx, http.StatusInternalServerError, "error to update value")
			return
		}
	}

	SendSucess(ctx, "operation-sync", contract)
}
