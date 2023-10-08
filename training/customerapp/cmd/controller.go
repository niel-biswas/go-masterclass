package main

import (
	"fmt"

	"customerapp/domain"
)

// CustomerController Organises the CRUD operations at UI layer.
type CustomerController struct {
	store domain.CustomerStore
}

// Add function to add new customer.
func (cc CustomerController) Add(c domain.Customer) {

	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Customer has been created\n")
}

// Update function to update a record.
func (cc CustomerController) Update(s string, c domain.Customer) {
	err := cc.store.Update(s, c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Customer %s record has been updated\n", s)

}

// Remove function to remove a record.
func (cc CustomerController) Remove(s string) {
	err := cc.store.Delete(s)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Customer %s record has been deleted\n", s)
}

// GetByCustomerId to get individual record based on id.
func (cc CustomerController) GetByCustomerId(s string) {
	c, err := cc.store.GetById(s)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("ID: %s\nName: %s\nE-mail:%s\n", c.ID, c.Name, c.Email)
}

// GetAll to get all customer records.
func (cc CustomerController) GetAll() {
	cs, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, c := range cs {
		fmt.Printf("ID: %s\nName: %s\nE-mail:%s\n", c.ID, c.Name, c.Email)
	}
}
