package controller_test

import (
	"bytes"
	"customerapp/controller"
	"customerapp/dbrepos"
	"customerapp/model"
	"customerapp/router"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TDD Based testing for Controller operations
func TestController(t *testing.T) {
	repo, _ := dbrepos.NewInmemoryRepository()
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	h := &controller.CustomerController{
		Repository: repo, // Injecting dependency
		Logger:     logger,
	}
	r := router.InitializeRoutes(h) // configure routes

	data := []model.Customer{}
	d := model.Customer{}
	var id string

	// Defining test cases and corresponding behaviours of controller actions
	testCases := []struct {
		name    string
		pattern string
		verb    string
		handle  func(t *testing.T)
	}{
		{
			name:    "Create endpoint using ResponseRecorder to create a Customer",
			pattern: "/api/customer",
			verb:    "POST",
			handle: func(t *testing.T) {
				var jsonStr = []byte(`{"name" : "Gopher", "email" : "workhard@gopher.com"}`)
				req, err := http.NewRequest("POST", "/api/customer", bytes.NewBuffer(jsonStr))
				if err != nil {
					t.Error(err)
				}
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				assert.Equal(t, http.StatusCreated, w.Code,
					fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusCreated, w.Code),
				)
			},
		},
		{
			name:    "GetAll endpoint using ResponseRecorder to fetch All Customers",
			pattern: "/api/customers",
			verb:    "GET",
			handle: func(t *testing.T) {
				req, err := http.NewRequest("GET", "/api/customers", nil)
				if err != nil {
					t.Error(err)
				}
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
					panic(err)
				}
				id = data[0].ID
				assert.Equal(t, http.StatusOK, w.Code,
					fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusOK, w.Code),
				)
			},
		},
		{
			name:    "Get endpoint using ResponseRecorder to fetch a Customer by id",
			pattern: "/api/customer/{id}",
			verb:    "GET",
			handle: func(t *testing.T) {
				req, err := http.NewRequest("GET", fmt.Sprintf("/api/customer/%s", id), nil)
				if err != nil {
					t.Error(err)
				}
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				assert.Equal(t, http.StatusOK, w.Code,
					fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusOK, w.Code),
				)
			},
		},
		{
			name:    "Put endpoint using ResponseRecorder to update a Customer by id",
			pattern: "/api/customer/{id}",
			verb:    "PUT",
			handle: func(t *testing.T) {
				var jsonStr = []byte(`{"name" : "GopherNew", "email" : "workhard@gophernew.com"}`)
				req, err := http.NewRequest("PUT", fmt.Sprintf("/api/customer/%s", id), bytes.NewBuffer(jsonStr))
				if err != nil {
					t.Error(err)
				}
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				assert.Equal(t, http.StatusNoContent, w.Code,
					fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusNoContent, w.Code),
				)
			},
		},
		{
			name:    "Get endpoint (Again - After Update) using ResponseRecorder to fetch a Customer by id",
			pattern: "/api/customer/{id}",
			verb:    "GET",
			handle: func(t *testing.T) {
				req, err := http.NewRequest("GET", fmt.Sprintf("/api/customer/%s", id), nil)
				if err != nil {
					t.Error(err)
				}
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				if err := json.Unmarshal(w.Body.Bytes(), &d); err != nil {
					panic(err)
				}
				if d.Name != "Gopher" && d.Email != "workhard@gopher.com" {
					logger.Info("Customer values updated successfully", zap.String("name", d.Name), zap.String("email", d.Email))
				} else {
					logger.Info("Customer values failed to update", zap.String("name", d.Name), zap.String("email", d.Email))
				}
				assert.Equal(t, http.StatusOK, w.Code,
					fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusOK, w.Code),
				)
			},
		},
		{
			name:    "Delete endpoint using ResponseRecorder to delete a Customer by id",
			pattern: "/api/customer/{id}",
			verb:    "DELETE",
			handle: func(t *testing.T) {
				req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/customer/%s", id), nil)
				if err != nil {
					t.Error(err)
				}
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				assert.Equal(t, http.StatusNoContent, w.Code,
					fmt.Sprintf("HTTP Status expected: %d, got: %d", http.StatusNoContent, w.Code),
				)
			},
		},
	}

	for _, tc := range testCases {
		logger.Info("Running test", zap.String("name", tc.name), zap.String("pattern", tc.pattern), zap.String("verb", tc.verb))
		t.Run(tc.name, tc.handle)
	}
}
