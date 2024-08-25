package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Server is running on http://localhost:3000")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(todos)
	})

	// * CREATE A TOOD
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := Todo{}

		if err := c.BodyParser(&todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, todo)

		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	// UPDATE A TODO
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(fiber.StatusOK).JSON(todos[i])
			}
		}

		return c.SendStatus(fiber.StatusNotFound)

	})

	log.Fatal(app.Listen(":3000"))
}
