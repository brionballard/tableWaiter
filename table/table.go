package table

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bxcodec/faker/v4"
)

const tableCount int = 12
const maxSeatCount int = 12

var tableSections = []string{"A", "B", "C", "D"}

const inMemoryTableDBFileName string = "tables.json"

// TableDB represents the table database.
type TableDB struct{}

// Table represents a single record in the table db
type Table struct {
	Id        string `json:"id"`
	Available bool   `json:"available"`
	Seats     int    `json:"seats"`
	Section   string `json:"section"`
}

// Init initializes the table database.
func (t *TableDB) Init() {
	writeInitialTableData(generateTableData())
}

// TableName returns the name of the table database.
func (t *TableDB) TableName() string {
	return inMemoryTableDBFileName
}

func generateTableData() []Table {
	tables := []Table{}

	for i := 1; i <= tableCount; i++ {
		table := Table{
			Id:        faker.UUIDHyphenated(),
			Available: true, // Every table is set to available when restaraunt is initialized
			Seats:     getRandomSeatCount(),
			Section:   getRandomSection(),
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
