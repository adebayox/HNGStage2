package config

import (
	"fmt"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db * gorm.DB
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("adebayox:my-secret-pw@tcp(localhost:3308)/some-mysql")

    db, err := gorm.Open("mysql", dsn)

    if err != nil {
        panic(err.Error())
    }

    return db
}

func GetDB() *gorm.DB{
	return db
}