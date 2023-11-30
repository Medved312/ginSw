package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/model"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	connectionString := "user=postgres password=1234 dbname=test sslmode=disable"
	var db, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Car{}, &model.Model{}, &model.Mark{}, &model.Transmission{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("Database in unavailable. Wait for %d sec./n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
