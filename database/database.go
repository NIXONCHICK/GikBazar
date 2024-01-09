package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:root@/goWeb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	return db, err
}
