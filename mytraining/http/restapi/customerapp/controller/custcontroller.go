package controller

import (
	"net/http"

	// internal
	"customerapp/model"

	// external
	"github.com/gorilla/mux"
)

type CustomerController struct {
	// Explicit dependency that hides dependent logic
	store model.Repository // CustomerStore value
	
}

func (ctl CustomerController) Post(w http.ResponseWriter, r *http.Request) {

}

func (ctl CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	var := mux.Vars(r)
}