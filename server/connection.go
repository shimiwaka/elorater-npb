package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func connectDB() *gorm.DB {
	s := Settings{}
	raw, err := os.ReadFile("./config/settings.json")
	if err != nil {
		return nil
	}

	err = json.Unmarshal(raw, &s)
	if err != nil {
		return nil
	}

	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.DB_username, s.DB_pass, s.DB_host, s.DB_port, s.DB_name)

	db, err := gorm.Open("mysql", connectQuery)

	if err != nil {
		return nil
	}

	return db
}
