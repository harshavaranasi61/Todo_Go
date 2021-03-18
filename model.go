package main

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `json:"Title"`
	Status string `json:"Status"`
}
