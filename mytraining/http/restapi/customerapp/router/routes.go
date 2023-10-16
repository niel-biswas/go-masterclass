package router

import (
	"customerapp/controller"
	"github.com/gorilla/mux"
)

func InitializeRoutes(h *controller.CustomerController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customers", h.GetAll).Methods("GET")
	r.HandleFunc("/api/customer/{id}", h.Get).Methods("GET")
	r.HandleFunc("/api/customer", h.Post).Methods("POST")
	r.HandleFunc("/api/customer/{id}", h.Put).Methods("PUT")
	r.HandleFunc("/api/customer/{id}", h.Delete).Methods("DELETE")
	return r
}
