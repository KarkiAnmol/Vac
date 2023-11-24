package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})
	err := app.Listen(":8081")
	if err != nil {
		panic(err)
	}
}
