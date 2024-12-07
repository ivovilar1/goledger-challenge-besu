package schemas

import (
	"fmt"
	"math/big"
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

// convert value to string to save into db
func (c *Contract) SetValue(value *big.Int) {
	c.Value = value.String()
}

// convert value to big int to use into app
func (c *Contract) GetValue() (*big.Int, error) {
	value := new(big.Int)
	if _, ok := value.SetString(c.Value, 10); !ok {
		return nil, fmt.Errorf("failed to convert string '%s' to *big.Int", c.Value)
	}
	return value, nil
}
