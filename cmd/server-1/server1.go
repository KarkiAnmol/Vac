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

//Event
type Event struct {
	ID                    string      `json:"id"`
	OccurredAt            string      `json:"occurredAt"`
	CorrelationID         string      `json:"correlationId"`
	PublishedAt           string      `json:"publishedAt"`
	SeriesID              int         `json:"seriesId"`
	SequenceNumber        int         `json:"sequenceNumber"`
	SessionSequenceNumber int         `json:"sessionSequenceNumber"`
	Events                []EventData `json:"events"`
}
type EventData struct {
	ID                string     `json:"id"`
	IncludesFullState bool       `json:"includesFullState"`
	Type              string     `json:"type"`
	Actor             ActorData  `json:"actor"`
	Action            string     `json:"action"`
	Target            TargetData `json:"target"`
	// SeriesStateDelta  SeriesDelta `json:"seriesStateDelta"`
	// SeriesState       SeriesState `json:"seriesState"`
}

// Actors
type ActorData struct {
	Type       string          `json:"type"`
	ID         string          `json:"id"`
	StateDelta ActorStateDelta `json:"stateDelta"`
	State      ActorState      `json:"state"`
}
type ActorStateDelta struct {
	ID   string        `json:"id"`
	Game ActorGameData `json:"game"`
}

type ActorGameData struct {
	ID          string `json:"id"`
	DamageDealt int    `json:"damageDealt"`
	// Add more fields as needed
}
type ActorState struct {
	ID     string        `json:"id"`
	TeamID string        `json:"teamId"`
	Side   string        `json:"side"`
	Series ActorSeries   `json:"series"`
	Game   ActorGameData `json:"game"`
	Name   string        `json:"name"`
}

// ActorSeries represents the "series" field inside ActorState
type ActorSeries struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	ParticipationStatus string         `json:"participationStatus"`
	Kills               int            `json:"kills"`
	KillAssistsReceived int            `json:"killAssistsReceived"`
	KillAssistsGiven    int            `json:"killAssistsGiven"`
	WeaponKills         map[string]int `json:"weaponKills"`
}

type TargetData struct {
	// Define fields specific to the TargetData
}

// type SeriesDelta struct {
// 	// Define fields specific to the SeriesDelta
// }

// type SeriesState struct {
// 	// Define fields specific to the SeriesState
// }

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

// func filter(event Event) (Event, error) {

// 	return event, nil
// }
