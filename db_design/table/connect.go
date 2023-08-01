package table

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type CustomDB struct {
	Db *gorm.DB
}

func (c *CustomDB) Insert(t any) error {
	return c.Db.WithContext(context.Background()).Create(t).Error
}

var dbInstance *CustomDB

func connect() error {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger.LogMode(logger.Info),
		//SkipDefaultTransaction: true,
	})
	//err = dbInstance.AutoMigrate(&table.User{}, &table.Contents{}, &table.Survey{})
	dbInstance = &CustomDB{db}
	return err
}

func init() {
	if err := connect(); err != nil {
		panic(err)
	}
}

func GetDB() *CustomDB {
	if dbInstance == nil {
		if err := connect(); err != nil {
			panic(err)
		}
	}
	return dbInstance
}
