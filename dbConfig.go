package main

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConfig() *gorm.DB {
	dsn := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASS") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DATABASE_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Todo{})
	return db
}
