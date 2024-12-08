package handler

import (
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivovilar1/goledger-challenge-besu/schemas"
)

func CheckHandler(ctx *gin.Context) {

	valueInBC, err := CallContract()
	if err != nil {
		logger.Errf("error to retrieve value: %v", err)
		SendError(ctx, http.StatusBadRequest, "error to retrieve value from blockchain")
		return
	}

	var valueInBCInt *big.Int
	switch v := valueInBC.(type) {
	case *big.Int:
		valueInBCInt = v
	default:
		valueInBCInt = new(big.Int)
		if _, ok := valueInBCInt.SetString(v.(string), 10); !ok {
			logger.Errf("failed to convert blockchain value '%v' to *big.Int", v)
			SendError(ctx, http.StatusBadRequest, "Invalid value format in blockchain")
			return
		}
	}

	var contract schemas.Contract
	if err := db.First(&contract).Error; err != nil {
		logger.Errf("error to retrieve value from database: %v", err)
		SendError(ctx, http.StatusInternalServerError, "error to retrieve value from database")
		return
	}

	convertedValue := new(big.Int)
	if _, ok := convertedValue.SetString(contract.Value, 10); !ok {
		logger.Errf("failed to convert string '%s' to *big.Int", contract.Value)
		SendError(ctx, http.StatusBadRequest, "Invalid value format in database")
		return
	}

	if valueInBCInt.Cmp(convertedValue) != 0 {
		SendSucess(ctx, "operation-check", false)
		return
	}

	SendSucess(ctx, "operation-check", true)
}
