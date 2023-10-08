// Package cmd entry point to customer record system.
package main

import (
	"fmt"

	"customerapp/domain"
	"customerapp/mapstore"
)

func main() {
	controller := CustomerController{ // initialize customer controller
		store: mapstore.NewMapStore(),
	}
	customer1 := domain.Customer{
		ID:    "cust101",
		Name:  "Rahul",
		Email: "rahul@gmail.com",
	}
	customer2 := domain.Customer{
		ID:    "cust102",
		Name:  "Shiju",
		Email: "shiju@gmail.com",
	}
	customer3 := domain.Customer{
		ID:    "cust102",
		Name:  "Raj",
		Email: "raj@gmail.com",
	}
	fmt.Println("***********************ADD************************************")
	controller.Add(customer1)
	controller.Add(customer2)
	fmt.Println("************************GETBYCustomerID************************")
	controller.GetByCustomerId("cust101")
	fmt.Println("**************************GETALL********************************")
	controller.GetAll()
	fmt.Println("**********************UPDATE**********************************")
	controller.Update("cust102", customer3)
	fmt.Println("*************************************************************")
	controller.GetByCustomerId("cust102")
	fmt.Println("*************************************************************")
	controller.Remove("cust102")
	fmt.Println("*************************************************************")
	controller.GetAll()
	fmt.Println("*************************************************************")
}
