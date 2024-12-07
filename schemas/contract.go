package schemas

import (
	"math/big"
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	Value *big.Int
}

type ContractResponse struct {
	ID        uint      `json:"id"`
	Value     *big.Int  `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
