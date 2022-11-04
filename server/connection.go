package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3307)/testdb?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	return db
}