package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"main/model"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	var db, err = gorm.Open("postgres", "user=postgres password=123 dbname=kinopoisk sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Movie{}, &model.Genre{}, &model.Movie_genre{})
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
