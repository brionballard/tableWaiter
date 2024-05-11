package table

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v4"
	"os"
	"tableWaiter/utils"
)

const tableCount int = 12
const maxSeatCount int = 12

var tableSections = []string{"A", "B", "C", "D"}

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
}

// Init initializes the table database.
func (t *TableDB) Init() {
	writeInitialTableData(generateTableData())
}

func generateTableData() []Table {
	tables := []Table{}

	for i := 1; i <= tableCount; i++ {
		reserved := utils.GenerateRandomBool()
		startTime := ""
		partyName := ""

		if reserved {
			startTime = utils.GenerateRandomTimeStringBetweenOpenAndClose()
			partyName = faker.LastName()
		}

		table := Table{
			Id:               faker.UUIDHyphenated(),
			Available:        reserved, // Every table is set to available when restaurant is initialized
			Seats:            getRandomSeatCount(),
			Section:          getRandomSection(),
			CurrentParty:     partyName,
			Reserved:         reserved,
			StartReservation: startTime,
			EndReservation:   "",
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

func getRandomSection() string {
	sectionNum, err := faker.RandomInt(0, len(tableSections)-1, 1)
	if err != nil {
		fmt.Println(err)
	}

	index := sectionNum[0]
	if sectionNum[0] > len(tableSections) {
		index = len(tableSections)
	}

	return tableSections[index]
}

func getRandomSeatCount() int {
	seatCount, err := faker.RandomInt(2, maxSeatCount, 1)
	if err != nil {
		fmt.Println(err)
	}

	return seatCount[0]
}
