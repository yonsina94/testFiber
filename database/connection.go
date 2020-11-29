package database

import (
	"github.com/yonsina94/testFiber/entities/person"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Conn active connection whit the database
var Conn *gorm.DB

func init() {
	if db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{}); err == nil {
		Conn = db
		db.AutoMigrate(&person.Person{})
	} else {
		panic(err)
	}
}
