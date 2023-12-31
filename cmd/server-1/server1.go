package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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
	SeriesID              string      `json:"seriesId"`
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

// ActorData represents the "actor" field in EventData
type ActorData struct {
	Type       string          `json:"type"`
	ID         string          `json:"id"`
	StateDelta ActorStateDelta `json:"stateDelta"`
	State      ActorState      `json:"state"`
}

// ActorStateDelta represents the "stateDelta" field inside ActorData
type ActorStateDelta struct {
	ID   string        `json:"id"`
	Game ActorGameData `json:"game"`
}

// ActorGameData represents the "game" field inside ActorStateDelta
type ActorGameData struct {
	ID          string `json:"id"`
	DamageDealt int    `json:"damageDealt"`
	// Add more fields as needed
}

// ActorState represents the "state" field inside ActorData
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

// TargetData represents the "target" field in EventData
type TargetData struct {
	Type       string           `json:"type"`
	ID         string           `json:"id"`
	StateDelta TargetStateDelta `json:"stateDelta"`
	State      TargetState      `json:"state"`
}

// TargetStateDelta represents the "stateDelta" field inside TargetData
type TargetStateDelta struct {
	ID   string         `json:"id"`
	Game TargetGameData `json:"game"`
	// Round TargetRoundData `json:"round"`
}

// TargetGameData represents the "game" field inside TargetStateDelta
type TargetGameData struct {
	ID            string `json:"id"`
	CurrentHealth int    `json:"currentHealth"`
	DamageTaken   int    `json:"damageTaken"`
	// CurrentArmor   int    `json:"currentArmor"`
	// Add more fields as needed
}

// TargetState represents the "state" field inside TargetData which is inside target field,which inturn is inside event field
type TargetState struct {
	ID     string         `json:"id"`
	TeamID string         `json:"teamId"`
	Side   string         `json:"side"`
	Series TargetSeries   `json:"series"`
	Game   TargetGameData `json:"game"`
	Name   string         `json:"name"`
}

// TargetSeries represents the "series" field inside TargetState
type TargetSeries struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	ParticipationStatus string         `json:"participationStatus"`
	Kills               int            `json:"kills"`
	KillAssistsReceived int            `json:"killAssistsReceived"`
	KillAssistsGiven    int            `json:"killAssistsGiven"`
	WeaponKills         map[string]int `json:"weaponKills"`
}

// type SeriesDelta struct {
// 	// Define fields specific to the SeriesDelta
// }

// type SeriesState struct {
// 	// Define fields specific to the SeriesState
// }

var e Event
var i = 0

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

		data := scanner.Text()
		// fmt.Printf("\n\nLength of Raw Data: %v\n", len(data))

		//Unmarshalling Raw JSON into struct filters the JSON
		err := json.Unmarshal([]byte(data), &e)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
			continue
		}
		// Filter function can be called here if needed
		// filteredData, err := filter(e)

		// fmt.Printf("\n\n\n%v : Filtered Data: \n%+v\n", count, e)
		// fmt.Printf("\n\nLength of Filtered Data: %v\n", int(reflect.TypeOf(e).Size()))

		// responseChan := make(chan error)
		for _, eventData := range e.Events {
			if i > 2 {
				i = 0
				// fmt.Println("Sleeping for 5 seconds...")
				time.Sleep(65 * time.Second)
			}
			i++
			// Check for events in priority order
			if eventData.Type == "team-completed-explodeBomb" {
				// printFilteredData("Team Completed Explode Bomb", e)
				// Call the function to send and receive data
				sendAndReceiveData(e)
			} else if eventData.Type == "team-won-round" {
				// printFilteredData("Team Won Round", e)
				sendAndReceiveData(e)
			} else if eventData.Type == "team-completed-defuseBomb" {
				// printFilteredData("Team Completed Defuse Bomb", e)

				sendAndReceiveData(e)
			} else if eventData.Type == "team-completed-beginDefuseWithoutKit" {
				// printFilteredData("Team Completed Begin Defuse Without Kit", e)
				sendAndReceiveData(e)
			} else if eventData.Type == "team-completed-plantBomb" {
				// printFilteredData("Team Completed Plant Bomb", e)
				sendAndReceiveData(e)
			} else if eventData.Type == "game-ended-round" {
				// printFilteredData("Game Ended Round", e)
				sendAndReceiveData(e)
			} else if eventData.Type == "game-started-round" {
				// printFilteredData("Game Started Round", e)
				sendAndReceiveData(e)
			} else if eventData.Type == "player-killed-player" {
				// printFilteredData("Player Killed Player", e)
				sendAndReceiveData(e)
			} else {
				continue
			}
			// // Wait for the response from the server before proceeding to the next iteration
			// if err := <-responseChan; err != nil {
			// 	log.Println("Error sending and receiving data:", err)
			// }
		}
		count++
		if count == 10000 {
			fmt.Println("Count:", count)
			break
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})
	log.Println("Server 1 listening on :8081")
	if err := app.Listen(":8081"); err != nil {
		panic(err)
	}

}

// func printFilteredData(eventType string, e Event) {
// 	x, err := json.Marshal(e)
// 	if err != nil {
// 		return
// 	}

func sendAndReceiveData(e Event) {
	serverURL := "http://localhost:8083/process"
	jsonData, err := json.Marshal(e)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}

	response, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending data:", err)
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Response from server: %s\n", responseBody)
}

// 	fmt.Printf("\nFiltered Data (%s): %+v", eventType, string(x))
// }

// func filter(event Event) (Event, error) {

// 	return event, nil
// }
