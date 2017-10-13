package main

import "fmt"

// Contact struct
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"`
	Phone string `json:"phone"`
}

// Contacts slice of contact
type Contacts []Contact

// String return the string of object
func (c *Contact) String() string {
	return fmt.Sprintf("%d, %s, %s, %s", c.ID, c.Name, c.City, c.Phone)
}
