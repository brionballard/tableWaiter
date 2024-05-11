package table

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v4"
	"log"
	"math/rand"
	"os"
	"tableWaiter/restaurant"
	"tableWaiter/utils"
	"time"
)

const inMemoryTableDBFileName string = "tables.json"

// TableDB represents the table database.
type TableDB struct{}

// Table represents a single record in the table db
type Table struct {
	Id                 string `json:"id"`
	Available          bool   `json:"available"`
	Seats              int    `json:"seats"`
	Section            string `json:"section"`
	CurrentParty       string `json:"currentParty"`
	PreviouslyReserved bool   `json:"previouslyReserved"`
	StartReservation   string `json:"startReservation"`
	EndReservation     string `json:"endReservation"`
	WaitTime           int    `json:"waitTime"`
}

// Init initializes the table database.
func (t *TableDB) Init() {
	writeInitialTableData(generateTableData())
}

// generateTableData generates fake table data for the resteraunt
func generateTableData() []Table {
	tables := []Table{}

	for i := 1; i <= restaurant.TableCount; i++ {
		reserved := utils.GenerateRandomBool()
		startTime := ""
		endTime := ""
		partyName := ""
		waitTime := 0

		if reserved {
			start, err := utils.GenerateRandomTimeString(restaurant.OpensAsInt, restaurant.ReservationsEndAsInt)
			if err != nil {
				log.Fatal("Failed to generate fake data.")
			}

			startTimeParsed, err := time.Parse(restaurant.TimeLayout, start)
			if err != nil {
				log.Fatal("Failed to parse startTime.")
			}

			// Generate a random duration between 30 minutes and 45 minutes
			randomMinutes := rand.Intn(16) + 30 // Generates a random integer between 30 and 45
			randomDuration := time.Duration(randomMinutes) * time.Minute
			end, err := utils.GenerateRandomTimeBetween(startTimeParsed, startTimeParsed.Add(randomDuration))
			if err != nil {
				log.Fatal("Failed to generate random time for end reservation.")
			}

			startTime = start
			endTime = end.Format(restaurant.TimeLayout)

			partyName = faker.LastName()
			// Calculate the wait time duration
			waitTimeDuration := end.Sub(startTimeParsed)

			// Convert the wait time duration to minutes
			waitTime = int(waitTimeDuration.Minutes())
		}

		table := Table{
			Id:                 faker.UUIDHyphenated(),
			Available:          reserved,
			Seats:              utils.GetRandomSBetweenMax(restaurant.MaxSeatingPerTable),
			Section:            utils.GetRandomStringFromSlice(restaurant.TableSections),
			CurrentParty:       partyName,
			PreviouslyReserved: reserved,
			StartReservation:   startTime,
			EndReservation:     endTime,
			WaitTime:           waitTime, // TODO:  difference in EndReservation - StartReservation
		}

		tables = append(tables, table)
	}

	return tables
}

// writeInitialTableData overwrites the data in the DB
// this function should only be used to initialize the database
func writeInitialTableData(t []Table) {
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}

	// Write data to the file
	err = os.WriteFile(inMemoryTableDBFileName, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func GetTables() ([]Table, error) {
	// Open the JSON file
	file, err := os.Open(inMemoryTableDBFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open tables file: %w", err)
	}
	defer file.Close()

	// Decode JSON into an array of Table structs
	var tables []Table
	if err := json.NewDecoder(file).Decode(&tables); err != nil {
		return nil, fmt.Errorf("failed to decode tables JSON: %w", err)
	}

	return tables, nil
}

func FilterTablesBySize(tables []Table, max int) []Table {
	var filteredTables []Table
	for _, table := range tables {
		if table.Seats >= max {
			filteredTables = append(filteredTables, table)
		}
	}
	return filteredTables
}

func FilterOutUnavailableTables(tables []Table) []Table {
	var available []Table
	for _, table := range tables {
		if table.Available {
			available = append(available, table)
		}
	}
	return available
}
