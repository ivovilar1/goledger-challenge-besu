package handler

import (
	"context"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	var result interface{}
	parsedAbi, err := abi.JSON(strings.NewReader(`[{
		"inputs": [],
		"name": "get",
		"outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
		"stateMutability": "view",
		"type": "function"
	}, {
		"inputs": [{"internalType": "uint256", "name": "x", "type": "uint256"}],
		"name": "set",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}]`))

	if err != nil {
		logger.Errf("error parsing abi: %v", err)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(goCtx, "http://localhost:8545")
	if err != nil {
		logger.Errf("error connecting to eth client: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress("0x42699A7612A82f1d9C36148af9C77354759b210b")
	caller := bind.CallOpts{
		Pending: false,
		Context: goCtx,
	}

	boundContract := bind.NewBoundContract(
		contractAddress,
		parsedAbi,
		client,
		client,
		client,
	)

	var output []interface{}
	err = boundContract.Call(&caller, &output, "get")
	if err != nil {
		logger.Errf("error calling contract: %v", err)
	}
	result = output

	SendSucess(ctx, "operation-set", result)
}
