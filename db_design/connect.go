package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dbInstance *gorm.DB

func connect() (err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := "guiwoo:guiwoo@tcp(127.0.0.1:3306)/guiwoo?charset=utf8mb4&parseTime=True&loc=Local"
	dbInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger.LogMode(logger.Info),
		//SkipDefaultTransaction: true,
	})
	//err = dbInstance.AutoMigrate(&table.User{}, &table.Contents{}, &table.Survey{})
	return err
}

func init() {
	if err := connect(); err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	if dbInstance == nil {
		if err := connect(); err != nil {
			panic(err)
		}
	}
	return dbInstance
}
