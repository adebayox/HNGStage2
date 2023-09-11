package config

import (
	"fmt"
	"os"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db * gorm.DB
)

func Connect() {
    dbHost := os.Getenv("DB_HOST")         
    dbPort := os.Getenv("DB_PORT")        
    dbUser := os.Getenv("DB_USER")        
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")        

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