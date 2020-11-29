package person

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"strconv"
)

type PersonController struct {
	service *PersonService
}

func (controller *PersonController) GetAllPersons(c *fiber.Ctx) error {
	persons := controller.service.GetAll()

	return c.Render("person/person.list", fiber.Map{
		"data": persons,
	})
}

func (controller *PersonController) GetPersonById(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 2, 10)
	person := controller.service.GetById(int(id))

	return c.Render("person/person.edit", fiber.Map{
		"data": person,
	})
}

func New(db *gorm.DB, router fiber.Router) *PersonController {
	service := Instance(db)

	controller := &PersonController{
		service: service,
	}

	router.Get("/", controller.GetAllPersons)
	router.Get("/:id", controller.GetPersonById)

	return controller
}
