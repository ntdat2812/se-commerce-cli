package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

type Config struct {
	User     string
	Password string
	Ip       string
	Port     string
	Database string
}



func Connect(config Config) error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, 
	config.Ip, config.Port, config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to dsn %v database: %v", dsn, err)
		return err
	}

	log.Println("Connecting to mysql database successfully")
	MysqlDB = db

	return nil
}
