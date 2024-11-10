package routes

import (
	"markeable/controllers"
	"markeable/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	app.Post("/receptionist/patient", middleware.AuthorizeRole("receptionist"), controllers.CreatePatient)
	app.Get("/receptionist/patients", middleware.AuthorizeRole("receptionist"), controllers.GetPatients)
	app.Put("/receptionist/patient/:id", middleware.AuthorizeRole("receptionist"), controllers.UpdatePatient)
	app.Delete("/receptionist/patient/:id", middleware.AuthorizeRole("receptionist"), controllers.DeletePatient)

	app.Post("/doctor/patient", middleware.AuthorizeRole("doctor"), controllers.CreatePatient)
	app.Get("/doctor/patients", middleware.AuthorizeRole("doctor"), controllers.GetPatients)
	app.Put("/doctor/patient/:id", middleware.AuthorizeRole("doctor"), controllers.UpdatePatient)
	app.Delete("/doctor/patient/:id", middleware.AuthorizeRole("doctor"), controllers.DeletePatient)
}
