package mysql_db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"reflect"
	"time"
)

var mysqlDB *MySqlDB

type MySqlDB struct {
	db *gorm.DB
}

func (m *MySqlDB) Create(value interface{}) error {
	return m.db.Create(value).Error
}

func (m *MySqlDB) Find(value interface{}) error {
	t := reflect.TypeOf(value)
	return m.db.Model(t).Take(value).Error
}

func connectDB() *gorm.DB {
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
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB() *MySqlDB {
	return mysqlDB
}

func Init() {
	if mysqlDB == nil {
		mysqlDB = &MySqlDB{
			db: connectDB(),
		}
	}
	return
}
