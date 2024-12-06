package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitDatabase() error {
	var err error

	db, err = InitSqlLite()
	if err != nil {
		return fmt.Errorf("error init sqlite: %v", err)
	}
	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
