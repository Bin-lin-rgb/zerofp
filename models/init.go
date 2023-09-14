package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"zerofp/define"
)

var DB *gorm.DB

func InitDB() {
	dsn := define.MysqlDSN + "/zerofp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}
	err = db.AutoMigrate(&User{}, &Product{}, &Device{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}
	DB = db
}
