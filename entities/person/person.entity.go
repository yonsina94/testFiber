package person

import (
	"time"

	"gorm.io/gorm"
)

//Person the entity of the table persons
type Person struct {
	gorm.Model

	//Name of the person
	Name string `gorm:"type:varchar(60);not null" json:"name" form:"name"`

	//SecondName of the person if exists
	SecondName string `gorm:"type:varchar(60)" json:"secondName" form:"secondName"`

	//FatherLastName parental lastname of the person
	FatherLastName string `gorm:"type:varchar(60);not null" json:"fatherLastName" form:"fatherLastName"`

	//MotherLastname maternal lastname of the person if exists
	MotherLastname string `gorm:"type:varchar(60)" json:"motherLastName" form:"motherLastName"`

	//BirtDate date of birt of the person
	BirtDate time.Time `gorm:"type:date" json:"birtdate" form:"birtdate"`

	//Email email address of the person
	Email string `gorm:"type:varchar(60)" json:"email" form:"email"`
}

func (p Person) Table() string {
	return "persons"
}
