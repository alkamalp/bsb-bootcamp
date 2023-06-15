package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

func Default() (*gorm.DB, error) {
	logLevel, _ := strconv.Atoi(os.Getenv("DB_LOG_MODE"))
	if logLevel == 0 {
		logLevel = 2
	}
	db, err := gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/smart-billing"),
		&gorm.Config{
			CreateBatchSize: 500,
			Logger:          logger.Default.LogMode(logger.LogLevel(logLevel)),
		})
	return db, err
}
