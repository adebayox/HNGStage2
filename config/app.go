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
    dbHost := "localhost"        
    dbPort := "3307"              
    dbUser := "root"             
    dbPassword := "my-secret-pw"   
    dbName := "personDB"           

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

    d, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }
    fmt.Println("Connected to MySQL database")
    db = d
}

func GetDB() *gorm.DB{
	return db
}