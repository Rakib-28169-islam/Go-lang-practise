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

    log.Fatal(app.Listen(":3000"))
}