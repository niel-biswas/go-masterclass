package controller

import (
	"encoding/json"
	"net/http"

	// internal
	"customerapp/model"

	// external
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CustomerController struct {
	// Explicit dependency that hides dependent logic
	Repository model.Repository // CustomerStore value
	Logger     *zap.Logger
}

// HTTP Post - /api/customer
func (ctl CustomerController) Post(w http.ResponseWriter, r *http.Request) {
	// Flushing any buffered log entries
	defer ctl.Logger.Sync()
	var customer model.Customer
	// Decode the incoming customer JSON
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a customer
	if err := ctl.Repository.Create(customer); err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)

		if err == model.ErrCustomerExists {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctl.Logger.Info("created a customer",
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusCreated)
}

// HTTP GetAll - /api/customers
func (ctl CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	// Flushing any buffered log entries
	defer ctl.Logger.Sync()
	// Get All
	if customers, err := ctl.Repository.GetAll(); err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		if err == model.ErrNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		j, err := json.Marshal(customers)
		if err != nil {
			ctl.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// HTTP GetAll - /api/customer/{id}
func (ctl CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	// Flushing any buffered log entries
	defer ctl.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	// Get by id
	if customers, err := ctl.Repository.GetById(id); err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		if err == model.ErrNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		j, err := json.Marshal(customers)
		if err != nil {
			ctl.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// HTTP Put - /api/customer/{id}
func (ctl CustomerController) Put(w http.ResponseWriter, r *http.Request) {
	// Flushing any buffered log entries
	defer ctl.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	var customer model.Customer
	// Decode the incoming JSON
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update a customer
	if err := ctl.Repository.Update(id, customer); err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctl.Logger.Info("updated customer",
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}

// HTTP Delete - /api/customer/{id}
func (ctl CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	// Flushing any buffered log entries
	defer ctl.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	// Delete a customer
	if err := ctl.Repository.Delete(id); err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctl.Logger.Info("deleted customer",
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}
