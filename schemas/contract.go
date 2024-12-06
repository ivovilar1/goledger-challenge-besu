package schemas

import (
	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	Value string
}
