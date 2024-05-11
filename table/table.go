package table

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v4"
	"log"
	"os"
	"tableWaiter/restaurant"
	"tableWaiter/utils"
)

const inMemoryTableDBFileName string = "tables.json"

// TableDB represents the table database.
type TableDB struct{}

// Table represents a single record in the table db
type Table struct {
	Id               string `json:"id"`
	Available        bool   `json:"available"`
	Seats            int    `json:"seats"`
	Section          string `json:"section"`
	CurrentParty     string `json:"currentParty"`
	Reserved         bool   `json:"reserved"`
	StartReservation string `json:"startReservation"`
	EndReservation   string `json:"endReservation"`
	WaitTime         int    `json:"waitTime"`
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
		partyName := ""
		waitTime := 0

		if reserved {
			timeString, err := utils.GenerateRandomTimeString(restaurant.OpensAsInt, restaurant.ReservationsEndAsInt)
			if err != nil {
				log.Fatal("Failed to generate fake data.")
			}
			startTime = timeString
			partyName = faker.LastName()
			waitTime = utils.GetRandomSBetweenMax(25)
		}

		table := Table{
			Id:               faker.UUIDHyphenated(),
			Available:        reserved, // Every table is set to available when restaurant is initialized
			Seats:            utils.GetRandomSBetweenMax(restaurant.MaxSeatingPerTable),
			Section:          utils.GetRandomStringFromSlice(restaurant.TableSections),
			CurrentParty:     partyName,
			Reserved:         reserved,
			StartReservation: startTime,
			EndReservation:   "",
			WaitTime:         waitTime,
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
