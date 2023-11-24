package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	jsonLines []string // jsonLines is a slice of strings
)

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
		fmt.Printf("%v : ", count)
		fmt.Println(scanner.Text())
		fmt.Println()
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})
	// app.Get("/get-json", func(c *fiber.Ctx) error {
	// 	// Read the JSON file
	// 	result, err := readJSONLinesFile()
	// 	// Respond with the entire JSON data
	// 	for _, line := range jsonLines {
	// 		fmt.Println(line)
	// 	}

	// 	return nil
	// })
	err = app.Listen(":8081")
	if err != nil {
		panic(err)
	}

}

// func readNextJSONLine() (string, error) {
// 	// For simplicity, this example reads a hard-coded JSONL string.
// 	input := "\"Hello\"\n\"World\"\n42\n"
// 	r := jsonl.NewReader(strings.NewReader(input))

// 	var result string

// 	// Read the next line
// 	err := r.ReadLines(func(data []byte) error {
// 		result = string(data)
// 		return nil
// 	})

// 	return result, err
// }
