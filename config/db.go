package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	InitEnv()
	var dbUser = os.Getenv("DB_USER")
	var dbPass = os.Getenv("DB_PASS")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbName = os.Getenv("DB_NAME")
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database not connected")
	}

	return db

}

func ConnDB() {

}
