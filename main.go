package main

import (
	"markeable/database"

	"markeable/routes"

	_ "markeable/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/gofiber/swagger"
)

//	@title			doctor patient site
//	@version		1.0
//	@description	Tis is a simple medical field.
//	@termsOfService	http://swagger.io/terms/
func main() {

	database.Connect()

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
