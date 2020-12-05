package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
	"github.com/yonsina94/testFiber/database"
	"github.com/yonsina94/testFiber/modules/person"

	"github.com/markbates/pkger"
)

func main() {
	engine := pug.NewFileSystem(pkger.Dir("/views"), ".pug")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	person.New(database.Conn, app)

	app.Listen(":8080")
}
