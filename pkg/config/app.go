package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error
	db, err = gorm.Open("mysql", "root:djdream3@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
