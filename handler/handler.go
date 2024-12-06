package handler

import (
	"github.com/ivovilar1/goledger-challenge-besu/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSqlite()
}
