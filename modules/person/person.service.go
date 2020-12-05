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

func (service *PersonService) Create(entity *entity.Person) error {
	return service.conn.Create(entity).Error
}

func (service *PersonService) Update(entity *entity.Person, id int) error {
	return service.conn.Save(entity).Error
}

func (service *PersonService) Delete(id int) error {
	return service.conn.Delete(&entity.Person{}, id).Error
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
