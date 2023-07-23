package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

type aa struct {
	Id   string `gorm:"primaryKey;column:id"`
	Name string `gorm:""`
}

func (a *aa) TableName() string {
	return "a"
}

type bb struct {
	AID  string `gorm:"primaryKey;column:a_id"`
	Name string `gorm:""`
}

func (b *bb) TableName() string {
	return "aa"
}

func TestDeadLock(t *testing.T) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := "root:rain45bow@tcp(127.0.0.1:3306)/guiwoo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger.LogMode(logger.Info),
	})

	var wg sync.WaitGroup
	wg.Add(2)

	var a aa
	var b bb

	start := make(chan int)

	go func() {
		defer wg.Done()
		<-start
		err := db.Transaction(func(tx *gorm.DB) error {
			err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&a).Error
			if err != nil {
				return err
			}
			time.Sleep(1 * time.Second)
			err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&b).Error
			return err
		})
		if err != nil {
			t.Errorf("db lock error err %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		<-start
		err := db.Transaction(func(tx *gorm.DB) error {
			err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&b).Error
			if err != nil {
				return err
			}
			time.Sleep(1 * time.Second)
			err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&a).Error
			return err
		})

		if err != nil {
			t.Errorf("db lock error err %v", err)
		}
	}()
	close(start)
	wg.Wait()
	fmt.Println("Test done")
}
