package main

// Contact struct
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"`
	Phone string `json:"phone"`
}

// Contacts slice of contact
type Contacts []Contact
