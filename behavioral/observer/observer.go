package observer

import "fmt"

type Observer interface {
	Update(msg string)
	GetID() string
}

type Customer struct {
	ID string
}

func (c *Customer) Update(msg string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.ID, msg)
}

func (c *Customer) GetID() string {
	return c.ID
}
