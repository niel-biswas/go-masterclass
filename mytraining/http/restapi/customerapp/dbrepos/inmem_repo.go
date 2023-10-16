package dbrepos

import (
	"errors"
	"time"

	// internal
	"customerapp/model"

	// external
	"github.com/gofrs/uuid"
)

// inmemoryRepository provides concrete implementation for the repository interface
type inmemoryRepository struct {
	custStore map[string]model.Customer
}

func NewInmemoryRepository() (model.Repository, error) {
	return &inmemoryRepository{
		custStore: make(map[string]model.Customer),
	}, nil
}

func (i *inmemoryRepository) isCustomerEmailExists(email string) bool {
	for _, v := range i.custStore {
		if v.Email == email {
			return true
		}
	}
	return false
}

func (i *inmemoryRepository) Create(c model.Customer) error {
	if _, ok := i.custStore[c.ID]; ok {
		return errors.New("Customer ID exists")
	}
	if i.isCustomerEmailExists(c.Email) {
		return model.ErrCustomerExists
	}
	c.CreatedOn = time.Now()
	// Create a Version 4 UUID.
	uid, _ := uuid.NewV4()
	c.ID = uid.String()
	i.custStore[c.ID] = c

	return nil
}

func (i *inmemoryRepository) Update(id string, c model.Customer) error {
	if _, ok := i.custStore[id]; !ok {
		return model.ErrNotFound
	}
	c.CreatedOn = time.Now()
	c.ID = id
	i.custStore[id] = c

	return nil
}

func (i *inmemoryRepository) Delete(id string) error {
	if _, ok := i.custStore[id]; !ok {
		return model.ErrNotFound
	}

	delete(i.custStore, id)

	return nil
}

func (i *inmemoryRepository) GetById(id string) (model.Customer, error) {
	if _, ok := i.custStore[id]; !ok {
		return model.Customer{}, model.ErrNotFound
	}

	return i.custStore[id], nil
}

func (i *inmemoryRepository) GetAll() ([]model.Customer, error) {
	if len(i.custStore) == 0 {
		return nil, model.ErrNotFound
	}
	customer := make([]model.Customer, 0, len(i.custStore))
	for _, v := range i.custStore {
		customer = append(customer, v)
	}

	return customer, nil
}
