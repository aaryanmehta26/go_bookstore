package models

import (
	"github.com/aaryanmehta26/go_bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name         string `gorm:""json:"name"`
	Author       string `json:"author"`
	Publications string `json:"publications"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
