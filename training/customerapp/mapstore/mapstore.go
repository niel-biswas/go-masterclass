// Package mapstore acts as backend mem data store.
package mapstore

import (
	"customerapp/domain"
)

// MapStore for memory based local data store.
type MapStore struct {
	store map[string]domain.Customer
}

// NewMapStore initialises the memory data store.
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

// Create inserts the record into mem datastore.
func (m *MapStore) Create(customer domain.Customer) error {
	if _, ok := m.store[customer.ID]; ok {
		return domain.ErrCustomerExists
	}
	m.store[customer.ID] = customer

	return nil
}

// Update updates the existing record into mem datastore.
func (m *MapStore) Update(s string, customer domain.Customer) error {
	if _, ok := m.store[s]; !ok {
		return domain.ErrCustomerNotExists
	}
	m.store[s] = customer

	return nil
}

// Delete deletes the record from mem datastore.
func (m *MapStore) Delete(s string) error {
	if _, ok := m.store[s]; !ok {
		return domain.ErrCustomerNotExists
	}

	delete(m.store, s)

	return nil
}

// GetById gets record based on id from mem datastore.
func (m *MapStore) GetById(s string) (domain.Customer, error) {
	if _, ok := m.store[s]; !ok {
		return domain.Customer{}, domain.ErrCustomerNotExists
	}

	return m.store[s], nil
}

// GetAll return all record from mem datastore.
func (m *MapStore) GetAll() ([]domain.Customer, error) {
	if len(m.store) == 0 {
		return nil, domain.ErrEmptyCustomers
	}
	customer := make([]domain.Customer, 0, len(m.store))
	for _, v := range m.store {
		customer = append(customer, v)
	}

	return customer, nil
}
