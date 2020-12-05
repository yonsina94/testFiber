package person

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yonsina94/testFiber/entities/person"
	"gorm.io/gorm"

	"fmt"
	"net/http"
	"strconv"
)

var routerPath string

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

func (controller *PersonController) PostNewPerson(c *fiber.Ctx) error {
	var person person.Person

	if err := c.BodyParser(&person); err == nil {
		if saverErr := controller.service.Create(&person); saverErr == nil {
			return c.Redirect(fmt.Sprintf("%s/", routerPath))
		} else {
			return c.Status(http.StatusBadRequest).JSONP(fiber.Map{"response": fiber.Map{"message": "occur a error when try to save the new person. Validate the person to add a try again !", "status": http.StatusBadRequest}}, "errorMessage")
		}
	} else {
		return c.Status(http.StatusBadRequest).JSONP(fiber.Map{
			"response": fiber.Map{
				"message": "The type of the received body not match whit the valid person type !",
				"status":  http.StatusBadRequest,
			},
		}, "errorMessage")
	}
}

func New(db *gorm.DB, app *fiber.App) *PersonController {
	service := Instance(db)

	controller := &PersonController{
		service: service,
	}

	routerPath = "/persons"

	app.Group(routerPath).Get("/", controller.GetAllPersons).Get("/:id", controller.GetPersonById).Post("/save", controller.PostNewPerson)

	return controller
}
