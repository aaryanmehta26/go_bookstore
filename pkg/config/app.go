package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// Replace these values with your actual database credentials and connection details
	connectionString := "username:password@tcp(hostname:port)/dbname?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
