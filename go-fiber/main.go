package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

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

	app.Put("/api/users/:id",func(c* fiber.Ctx)error{
		id := c.Params("id")
		return c.JSON(fiber.Map{
			"message": "User " + id + " created successfully",
		})
	})

    log.Fatal(app.Listen(":3000"))
}