package models

import (
	"github.com/hng/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Person struct{
	gorm.Model
	Name string `gorm:""json:"name"`
}

func init() {
	config.Connect()
	db = config .GetDB()
	db.AutoMigrate(&Person{})
}

func GetDB() *gorm.DB {
    return db
}

func (b *Person) CreatePerson() *Person{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetPersonById(Id int64) (*Person, *gorm.DB){
	var getPerson Person
	db := db.Where("ID=?", Id).Find(&getPerson)
	return &getPerson, db
}

func DeletePerson(ID int64) Person{
	var person Person
	db.Where("ID=?", ID).Delete(person)
	return person
}

func DeletePersonByName(name string) Person {
    var person Person
    db.Where("name = ?", name).Delete(&person)
    return person
}


func GetPersonByName(name string) (*Person, *gorm.DB) {
    var getPerson Person
    db := db.Where("name = ?", name).First(&getPerson)
    return &getPerson, db
}
