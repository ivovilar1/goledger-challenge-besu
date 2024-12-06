package config

import (
	"os"

	"github.com/ivovilar1/goledger-challenge-besu/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//check if database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database not found, creating...")
		// create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("sqlite open error: %v", err)
		return nil, err
	}

	// migrate schema
	err = db.AutoMigrate(&schemas.Contract{})
	if err != nil {
		logger.Errf("sqlite automigration error: %v", err)
	}

	return db, nil
}
