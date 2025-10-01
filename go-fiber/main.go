package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber in VSCode!")
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"users": []string{"John", "Jane", "Bob"},
		})
	})

	// Dynamic route with parameter
	app.Put("/api/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{
			"message": "User " + id + " created successfully",
		})
	})
	//query parameter
	app.Get("/api/users/search", func(c *fiber.Ctx) error {
		query := c.Query("q")
		page := c.Query("page", "1") //default value 1
		return c.JSON(fiber.Map{
			"query": query,
			"page":  page,
		})
	})
	// Multiple parameters
	app.Get("/user/:id/post/:postId", func(c *fiber.Ctx) error {
		id := c.Params("id")
		postId := c.Params("postId")
		return c.JSON(fiber.Map{
			"userId": id,
			"postId": postId,
		})
	})

	app.Post("/api/users/create-user", func(c *fiber.Ctx) error {

		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  400,
				"error":   "Invalid request body",
				"details": err.Error(),
			})
		}
		log.Printf("Received user: %+v\n", user)
		fmt.Printf("%+v\n", user)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User created successfully",
			"user":    user,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
