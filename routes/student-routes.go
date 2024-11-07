package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lgarg12/Fealty_Assignment/handlers"
)

func Routers(app *fiber.App) {
	app.Get("/students", handlers.GetAllStudents)
	app.Post("/students", handlers.CreateStudent)
	app.Get("/students/:id", handlers.GetStudentById)
	app.Delete("/students/:id", handlers.DeleteStudentById)
}
