package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SetHandler(ctx *gin.Context) {

	request := SetContractValueRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	value := request.Value

	abi, err := abi.JSON(strings.NewReader(`[{
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
	}]`)) // found under besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json

	if err != nil {
		logger.Err(err)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(goCtx, "http://localhost:8545") // e.g., http://localhost:8545
	if err != nil {
		logger.Errf("error dialing node: %v", err)
	}

	logger.Info("querying chain id")

	chainId, err := client.ChainID(goCtx)
	if err != nil {
		logger.Errf("error querying chain id: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress("0x42699A7612A82f1d9C36148af9C77354759b210b") // will be returned during startDev.sh execution

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	priv, err := crypto.HexToECDSA("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63") // this can be found in the genesis.json file
	if err != nil {
		logger.Errf("error loading private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainId)
	if err != nil {
		logger.Errf("error creating transactor: %v", err)
	}

	tx, err := boundContract.Transact(auth, "set", value)
	if err != nil {
		logger.Errf("error transacting: %v", err)
	}

	if err != nil {
		logger.Errf("Error during transaction: %v", err)
		SendError(ctx, http.StatusInternalServerError, "Transaction failed")
		return
	}

	receipt, err := bind.WaitMined(
		context.Background(),
		client,
		tx,
	)
	if err != nil {
		logger.Errf("error waiting for transaction to be mined: %v", err)
	}

	SendSucess(ctx, "operation-set", receipt)
}
