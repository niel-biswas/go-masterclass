// Package domain acts as customer data model
package domain

import "errors"

var (
	ErrCustomerExists    error = errors.New("customer already exists")
	ErrCustomerNotExists error = errors.New("customer doesn't exist")
	ErrEmptyCustomers    error = errors.New("customers are empty")
)

// Customer data model.
type Customer struct {
	ID    string
	Name  string
	Email string
}

// CustomerStore interface for CRUD operation in backend.
type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}
