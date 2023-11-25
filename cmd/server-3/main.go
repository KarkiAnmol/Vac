package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	Data string `json:"data"`
}

type ResponseData struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	app := fiber.New()

	// Define a route to handle incoming JSON data
	app.Post("/process", func(c *fiber.Ctx) error {
		// Parse the JSON data from the request
		var requestData RequestData
		if err := c.BodyParser(&requestData); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		// Process the data (for demonstration, just add a timestamp)
		responseData := ResponseData{
			Message:   fmt.Sprintf("Received: %s", requestData.Data),
			Timestamp: time.Now(),
		}

		// Send the processed data back as the response
		return c.JSON(responseData)
	})

	// Start the server on port 8082
	if err := app.Listen(":8082"); err != nil {
		log.Fatal(err)
	}
}
