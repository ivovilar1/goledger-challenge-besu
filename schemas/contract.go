package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	Value string
}

type ContractResponse struct {
	ID        uint      `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
