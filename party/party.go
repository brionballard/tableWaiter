package party

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"os"
	"tableWaiter/table"
)

type Party struct {
	Name  string `json:"name"`
	Size  int    `json:"size"`
	Phone string `json:"phone"`
}

// AskForInfo asks for important information to seat the party
// It returns an error if there is an issue reading the input.
func (p *Party) AskForInfo() error {
	if err := p.AskForName(); err != nil {
		return err
	}
	if err := p.AskForSize(); err != nil {
		return err
	}
	if err := p.AskForPhone(); err != nil {
		return err
	}
	return nil
}

// AskForName asks the user for the name of the party and assigns it to the Party struct.
// It returns an error if there is an issue reading the input.
func (p *Party) AskForName() error {
	fmt.Println("Can I get the name for your party?")
	_, err := fmt.Scanln(&p.Name)
	if err != nil {
		return fmt.Errorf("failed to gather party name: %w", err)
	}

	return nil
}

// AskForSize asks the user for the size of the party and assigns it to the Party struct.
// It returns an error if there is an issue reading the input.
func (p *Party) AskForSize() error {
	fmt.Println("And how many are in your party?")
	_, err := fmt.Scanln(&p.Size)
	if err != nil {
		return fmt.Errorf("failed to gather party size: %w", err)
	}

	return nil
}

// AskForPhone asks the user for the phone number of the party and assigns it to the Party struct.
// It returns an error if there is an issue reading the input.
func (p *Party) AskForPhone() error {
	fmt.Println("May I have a phone number for your party?")
	_, err := fmt.Scanln(&p.Phone)
	if err != nil {
		return fmt.Errorf("failed to gather party phone: %w", err)
	}

	return nil
}

func (p *Party) Seat() {
	tables, err := table.GetTables()
	if err != nil {
		_ = fmt.Errorf("cannot currently retrieve tables: %w", err)
	}

	tablesWithAdequateSeating := table.FilterTablesBySize(tables, p.Size)
	availableTables := table.FilterOutUnavailableTables(tablesWithAdequateSeating)

	fmt.Println(len(availableTables))

	if len(availableTables) > 0 {
		// pick a random table and seat the party
		tableIndex, _ := faker.RandomInt(1, len(availableTables)-1)
		selectedTable := availableTables[tableIndex[0]]

		p.FollowMe(selectedTable)
	}

	// Get average wait time of available tables
	// inform party there is an x minute wait to be seated
	// do some sort of set timeout type function while printing hashes
	// select a table,
	//	update CurrentParty,
	//	update Available to false
	//  update StartReservation
	//	update EndReservation to 1.5 hours
}

func (p *Party) FollowMe(tbl table.Table) {
	fmt.Printf("Okay %s you can follow me to table %s in section %s \n", p.Name, tbl.Id, tbl.Section)
	os.Exit(0)
}
