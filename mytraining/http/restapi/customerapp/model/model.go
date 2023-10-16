package model

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("No records found")
var ErrCustomerExists = errors.New("Customer already exists")

type Customer struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	CreatedOn time.Time `json:"createdon,omitempty" bson:"createdon,omitempty"`
}

type Repository interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}
