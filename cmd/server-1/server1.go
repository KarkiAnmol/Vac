package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	jsonLines []string // jsonLines is a slice of strings
)

type Event struct {
	ID            string `json:"id"`
	OccurredAt    string `json:"occurredAt"`
	CorrelationID string `json:"correlationId"`
}

var e Event

func main() {
	app := fiber.New()

	file, err := os.Open("2579048_events.jsonl")
	if err != nil {
		log.Fatal("File doesnot exist")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCapacity = 10 * 1024 * 1024 // Set a larger buffer size (e.g., 10 MB)
	// Increase buffer size to handle longer lines
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	count := 1
	for scanner.Scan() {
		fmt.Println()
		fmt.Printf("%v : %s ", count, scanner.Text())
		data := scanner.Text()
		//Unmarshalling Raw JSON into struct filters the JSON
		err := json.Unmarshal([]byte(data), &e)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
			continue
		}
		// Filter function can be called here if needed
		// filteredData, err := filter(e)

		fmt.Printf("\nFiltered Data: \n%+v\n", e)
		// later-on pass this filtered data to server 3 where the commentary will be generated

		count++
		break
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})

	if err := app.Listen(":8081"); err != nil {
		panic(err)
	}

}
func filter(event Event) (Event, error) {

	return event, nil
}
