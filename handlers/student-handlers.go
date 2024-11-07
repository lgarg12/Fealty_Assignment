package handlers

import (
	"strconv"


	"github.com/gofiber/fiber/v2"
	"github.com/lgarg12/Fealty_Assignment/models"
)

func GetAllStudents(c *fiber.Ctx) error {
	var students []models.Student
	for _, student := range models.Students {
		students = append(students, student)
	}

	return c.JSON(students)
}

func CreateStudent(c *fiber.Ctx) error {
	student := new(models.Student)

	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Input"})
	}

	newId := len(models.Students) + 1

	student.Id = newId
	models.Students[newId] = *student

	return c.Status(fiber.StatusCreated).JSON(student)

}

func GetStudentById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	student, exists := models.Students[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	return c.JSON(student)
}

func DeleteStudentById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	_, exists := models.Students[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	delete(models.Students, id)
	return c.SendStatus(fiber.StatusNoContent)
}


