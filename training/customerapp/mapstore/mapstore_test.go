package mapstore_test

import (
	"errors"
	"reflect"
	"slices"
	"testing"

	"customerapp/domain"
	"customerapp/mapstore"
)

func TestMapStore_Create(t *testing.T) {

	m := mapstore.NewMapStore()
	tests := []struct {
		name     string
		customer domain.Customer
		wantErr  error
	}{
		{
			name: "Add customer",
			customer: domain.Customer{
				ID:    "cust101",
				Name:  "Rahul",
				Email: "rahul@gmail.com",
			},
			wantErr: nil,
		},
		{
			name: "adding same customer",
			customer: domain.Customer{
				ID:    "cust101",
				Name:  "Rahul",
				Email: "rahul@gmail.com",
			},
			wantErr: domain.ErrCustomerExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.Create(tt.customer)
			if err != nil && tt.wantErr == nil {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr != nil {
				if errors.Is(err, tt.wantErr) == false {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestMapStore_Delete(t *testing.T) {
	m := mapstore.NewMapStore()
	customer := domain.Customer{
		ID:    "cust101",
		Name:  "Rahul",
		Email: "rahul@gmail.com",
	}
	_ = m.Create(customer)
	tests := []struct {
		name     string
		customer string
		wantErr  error
	}{
		{
			name:     "delete customer",
			customer: "cust101",
			wantErr:  nil,
		},
		{
			name:     "delete same customer again",
			customer: "cust101",
			wantErr:  domain.ErrCustomerNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.Delete(tt.customer)
			if err != nil && tt.wantErr == nil {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != nil {
				if errors.Is(err, tt.wantErr) == false {
					t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestMapStore_GetAll(t *testing.T) {
	m := mapstore.NewMapStore()
	customer := domain.Customer{
		ID:    "cust101",
		Name:  "Rahul",
		Email: "rahul@gmail.com",
	}
	_ = m.Create(customer)
	tests := []struct {
		name    string
		want    []domain.Customer
		wantErr error
	}{
		{
			name: "getting all customer",
			want: []domain.Customer{
				domain.Customer{
					ID:    "cust101",
					Name:  "Rahul",
					Email: "rahul@gmail.com",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetAll()
			if err != nil && tt.wantErr == nil {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			//}
			if !slices.Equal(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStore_GetById(t *testing.T) {
	m := mapstore.NewMapStore()
	customer := domain.Customer{
		ID:    "cust101",
		Name:  "Rahul",
		Email: "rahul@gmail.com",
	}
	_ = m.Create(customer)
	tests := []struct {
		name    string
		id      string
		want    domain.Customer
		wantErr error
	}{
		{
			name: "getting by customer id",
			id:   "cust101",
			want: domain.Customer{
				ID:    "cust101",
				Name:  "Rahul",
				Email: "rahul@gmail.com",
			},

			wantErr: nil,
		},
		{
			name:    "getting by non existing  customer id",
			id:      "cust102",
			want:    domain.Customer{},
			wantErr: domain.ErrCustomerNotExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetById(tt.id)
			if err != nil && tt.wantErr == nil {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != nil {
				if errors.Is(err, tt.wantErr) == false {
					t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStore_Update(t *testing.T) {
	m := mapstore.NewMapStore()
	customer := domain.Customer{
		ID:    "cust101",
		Name:  "Rahul",
		Email: "rahul@gmail.com",
	}
	_ = m.Create(customer)
	tests := []struct {
		name     string
		id       string
		customer domain.Customer
		wantErr  error
	}{
		{
			name: "update customer",
			id:   "cust101",
			customer: domain.Customer{
				ID:    "cust101",
				Name:  "Rahul Krishnan",
				Email: "rahulk@gmail.com",
			},
			wantErr: nil,
		},
		{
			name:    "update non existing customer",
			id:      "cust102",
			wantErr: domain.ErrCustomerNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.Update(tt.id, tt.customer)
			if err != nil && tt.wantErr == nil {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != nil {
				if errors.Is(err, tt.wantErr) == false {
					t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestNewMapStore(t *testing.T) {
	tests := []struct {
		name string
		want *mapstore.MapStore
	}{
		{
			name: "check creation of map store",
			want: mapstore.NewMapStore(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapstore.NewMapStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMapStore() = %v, want %v", got, tt.want)
			}
		})
	}
}
