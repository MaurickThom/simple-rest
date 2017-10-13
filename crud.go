package main

var listado Contacts

// loadInitialData
func loadInitialData() {
	c1 := Contact{
		ID:    1,
		Name:  "LEIDY",
		City:  "BOGOTA",
		Phone: "444444",
	}
	c2 := Contact{
		ID:    2,
		Name:  "CAROL",
		City:  "BOGOTA",
		Phone: "555555",
	}
	c3 := Contact{
		ID:    3,
		Name:  "JUAN",
		City:  "CALI",
		Phone: "6666666",
	}
	listado = append(listado, &c1, &c2, &c3)
}

// getAll returns all contacts
func getAll() Contacts {
	return listado
}

// getByID return a contact by id
func getByID(id int) *Contact {
	for _, v := range listado {
		if v.ID == id {
			return v
		}
	}

	return nil
}

// add insert a contact
func add(c *Contact) {
	c.ID = getMaxID()
	listado = append(listado, c)
}

// getMaxID from listado
func getMaxID() int {
	size := len(listado)
	if size > 0 {
		return listado[size-1].ID + 1
	}

	return 1
}

// update
func update(id int, c *Contact) {
	for _, v := range listado {
		if v.ID == id {
			v = c
			return
		}
	}
}

// delete
func delete(id int) {
	listado = append(listado[:id-1], listado[id:]...)
}
