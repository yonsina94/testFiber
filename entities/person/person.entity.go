package person

import (
	"time"

	"gorm.io/gorm"
)

//Person the entity of the table persons
type Person struct {
	gorm.Model

	//Name of the person
	Name string `gorm:"type:varchar(60);not null"`

	//SecondName of the person if exists
	SecondName string `gorm:"type:varchar(60)"`

	//FatherLastName parental lastname of the person
	FatherLastName string `gorm:"type:varchar(60);not null"`

	//MotherLastname maternal lastname of the person if exists
	MotherLastname string `gorm:"type:varchar(60)"`

	//BirtDate date of birt of the person
	BirtDate time.Time `gorm:"type:date"`

	//Email email address of the person
	Email string `gorm:"type:varchar(60)"`
}

func (p Person) Table() string {
	return "persons"
}
