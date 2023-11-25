// server3.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

var eventData Event

func main() {
	app := fiber.New()

	// Define a route to handle incoming POST requests on /process
	app.Post("/process", func(c *fiber.Ctx) error {
		// Parse the JSON data from the request body
		if err := c.BodyParser(&eventData); err != nil {
			log.Println("Error parsing JSON:", err)
			return c.SendStatus(http.StatusBadRequest)
		}

		// Process the event data as needed, i.e., write code for generating commentary
		// For now, just print the received data
		fmt.Printf("Received Event Data: %+v\n", eventData)

		// Send a simple response
		response := "Data processed successfully"
		fmt.Printf("Response to client: %s\n", response)
		return c.SendString(response)
	})

	// Define a route for a simple ping
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})
	log.Println("Server 3 listening on port 8083")
	// Start the server on port 8083
	if err := app.Listen(":8083"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}