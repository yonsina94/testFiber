package person

import (
	"gorm.io/gorm"

	entity "github.com/yonsina94/testFiber/entities/person"
)

type PersonService struct {
	conn *gorm.DB
}

func Instance(conn *gorm.DB) *PersonService {
	return &PersonService{
		conn: conn,
	}
}

func (service *PersonService) Create(entity *entity.Person) *entity.Person {
	service.conn.Create(entity)

	return entity
}

func (service *PersonService) Update(entity *entity.Person, id int) *entity.Person {
	service.conn.Save(entity)

	return entity
}

func (service *PersonService) Delete(id int) bool {
	if err := service.conn.Delete(&entity.Person{}, id).Error; err == nil {
		return true
	} else {
		return false
	}
}

func (service *PersonService) GetAll() []entity.Person {
	var persons []entity.Person

	service.conn.Find(&persons)

	return persons
}

func (service *PersonService) GetById(id int) entity.Person {
	var person entity.Person

	service.conn.First(&person, id)

	return person
}
