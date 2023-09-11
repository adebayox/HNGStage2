package config

import (
	"fmt"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db * gorm.DB
)

func Connect() {
	dsn := "root:my-secret-pw@tcp(localhost:3307)/personDB?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to sql database")
	db = d
}

func GetDB() *gorm.DB{
	return db
}