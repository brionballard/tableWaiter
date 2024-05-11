package party

import "fmt"

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
	// read in table data
	// find table with minimum amount of seats
	// if available reserve
	// else if no available tables -- set waiting time
	// in x time find available table
	// Set currentParty to table
}
