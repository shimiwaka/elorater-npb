package main

import (
	"encoding/json"
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	s := Settings{}
	raw, err := os.ReadFile("./config/settings.json")
	if err != nil {
		panic("failed to open config file")
	}

	json.Unmarshal(raw, &s)

	sqlConnect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
					s.DB_username, s.DB_pass, s.DB_host, s.DB_port, s.DB_name)

	db, err := gorm.Open("mysql", sqlConnect)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}