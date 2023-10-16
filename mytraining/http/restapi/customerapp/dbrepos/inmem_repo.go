package dbrepos

import (
	"customerapp/model"
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

func (i *inmemoryRepository) Create(c model.Customer) error {
	if _, ok := i.custStore[c.ID]; ok {
		return model.ErrCustomerExists
	}
	i.custStore[c.ID] = c

	return nil
}

func (i *inmemoryRepository) Update(id string, c model.Customer) error {
	if _, ok := i.custStore[id]; !ok {
		return model.ErrNotFound
	}
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
