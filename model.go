package main

import "fmt"

// Contact struct
type Contact struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	City        string       `json:"city"`
	Phone       string       `json:"phone"`
	Navigations []Navigation `json:"navigation"`
}

// Navigation struct
type Navigation struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

// Contacts slice of contact
type Contacts []*Contact

// String return the string of object
func (c *Contact) String() string {
	return fmt.Sprintf("%d, %s, %s, %s", c.ID, c.Name, c.City, c.Phone)
}

// String of contacts
func (cs *Contacts) String() string {
	var r string
	for _, c := range *cs {
		r += fmt.Sprintf("%d, %s, %s, %s\n", c.ID, c.Name, c.City, c.Phone)
	}
	return r
}
