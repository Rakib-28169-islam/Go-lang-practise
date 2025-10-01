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

	// Dynamic route with parameter
	app.Put("/api/users/:id",func(c* fiber.Ctx)error{
		id := c.Params("id")
		return c.JSON(fiber.Map{
			"message": "User " + id + " created successfully",
		})
	})
	//query parameter
	app.Get("/api/users/search",func (c *fiber.Ctx)error{
		query :=c.Query("q")
		page := c.Query("page","1")//default value 1
		return c.JSON(fiber.Map{
			"query": query,
			"page": page,
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

    log.Fatal(app.Listen(":3000"))
}