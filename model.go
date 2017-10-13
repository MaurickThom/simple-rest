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
type Contacts []*Contact

// String return the string of object
func (c *Contact) String() string {
	return fmt.Sprintf("%d, %s, %s, %s", c.ID, c.Name, c.City, c.Phone)
}

// String of contacts
func (cs *Contacts) String() string {
	var r string
	for _, c := range *cs {
		fmt.Printf("%#v", c)
		r += fmt.Sprintf("%d, %s, %s, %s\n", c.ID, c.Name, c.City, c.Phone)
	}
	fmt.Println(r)
	return r
}
