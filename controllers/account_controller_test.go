package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Elagoht/go-pass/db"
	"github.com/Elagoht/go-pass/models"
	"github.com/gorilla/mux"
)

// setupTestDB initializes a test database
func setupTestDB() {
	// Use a test database file
	os.Setenv("DB_PATH", "./test_accounts.db")
	db.InitDB()
}

// cleanupTestDB removes the test database file
func cleanupTestDB() {
	os.Remove("./test_accounts.db")
}

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	// Cleanup
	cleanupTestDB()
	os.Exit(code)
}

type TestCreatePayload struct {
	name       string
	payload    models.Account
	wantStatus int
}

type TestIdGetter struct {
	name       string
	id         string
	wantStatus int
}

type TestUpdatePayload struct {
	name       string
	id         string
	payload    models.Account
	wantStatus int
}

/* Successful and missing fields tests for account creation */
func TestCreateAccount(tester *testing.T) {
	setupTestDB()
	controller := NewAccountController()

	tests := []TestCreatePayload{
		{ // Valid Account
			name: "Valid Account",
			payload: models.Account{
				Platform:   "GitHub",
				URL:        "https://github.com",
				Identity:   "testuser",
				Passphrase: "testpass123",
				Notes:      "Test account",
			},
			wantStatus: http.StatusCreated,
		},
		{ // Invalid Account - Missing Fields
			name: "Invalid Account - Missing Fields",
			payload: models.Account{
				Platform: "GitHub",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	// Iterate through the tests and run them
	for _, testValue := range tests {
		tester.Run(testValue.name, func(test *testing.T) {
			// Convert the payload to JSON bytes
			body, _ := json.Marshal(testValue.payload)

			// Create and record a new request and record the response
			request := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
			recorder := httptest.NewRecorder()
			controller.CreateAccount(recorder, request)

			if recorder.Code != testValue.wantStatus {
				test.Errorf("CreateAccount() status = %v, want %v", recorder.Code, testValue.wantStatus)
			}
		})
	}
}

/* Test the GetAccounts function */
func TestGetAccounts(tester *testing.T) {
	setupTestDB()
	controller := NewAccountController()

	// Create and record a new request and record the response
	request := httptest.NewRequest("GET", "/accounts", nil)
	recorder := httptest.NewRecorder()
	controller.GetAccounts(recorder, request)

	if recorder.Code != http.StatusOK {
		tester.Errorf("GetAccounts() status = %v, want %v", recorder.Code, http.StatusOK)
	}
}

/* Test the GetAccount function */
func TestGetAccount(tester *testing.T) {
	setupTestDB()
	controller := NewAccountController()

	tests := []TestIdGetter{
		{
			name:       "Non-existent Account",
			id:         "999",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, testValue := range tests {
		tester.Run(testValue.name, func(test *testing.T) {
			request := httptest.NewRequest("GET", "/accounts/"+testValue.id, nil)
			recorder := httptest.NewRecorder()

			// Set up the router with the ID parameter
			router := mux.NewRouter()
			router.HandleFunc("/accounts/{id}", controller.GetAccount).Methods("GET")
			router.ServeHTTP(recorder, request)

			if recorder.Code != testValue.wantStatus {
				test.Errorf("GetAccount() status = %v, want %v", recorder.Code, testValue.wantStatus)
			}
		})
	}
}

/* Test the UpdateAccount function */
func TestUpdateAccount(tester *testing.T) {
	setupTestDB()
	controller := NewAccountController()

	tests := []TestUpdatePayload{
		{
			name: "Update Non-existent Account",
			id:   "999",
			payload: models.Account{
				Platform:   "GitHub",
				URL:        "https://github.com",
				Identity:   "updateduser",
				Passphrase: "updatedpass123",
				Notes:      "Updated test account",
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, testValue := range tests {
		tester.Run(testValue.name, func(test *testing.T) {
			body, _ := json.Marshal(testValue.payload)
			request := httptest.NewRequest("PUT", "/accounts/"+testValue.id, bytes.NewBuffer(body))
			recorder := httptest.NewRecorder()

			// Set up the router with the ID parameter
			router := mux.NewRouter()
			router.HandleFunc("/accounts/{id}", controller.UpdateAccount).Methods("PUT")
			router.ServeHTTP(recorder, request)

			if recorder.Code != testValue.wantStatus {
				test.Errorf("UpdateAccount() status = %v, want %v", recorder.Code, testValue.wantStatus)
			}
		})
	}
}

/* Test the DeleteAccount function */
func TestDeleteAccount(tester *testing.T) {
	setupTestDB()
	controller := NewAccountController()

	tests := []TestIdGetter{
		{
			name:       "Delete Non-existent Account",
			id:         "999",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, testValue := range tests {
		tester.Run(testValue.name, func(test *testing.T) {
			request := httptest.NewRequest("DELETE", "/accounts/"+testValue.id, nil)
			recorder := httptest.NewRecorder()

			// Set up the router with the ID parameter
			router := mux.NewRouter()
			router.HandleFunc("/accounts/{id}", controller.DeleteAccount).Methods("DELETE")
			router.ServeHTTP(recorder, request)

			if recorder.Code != testValue.wantStatus {
				test.Errorf("DeleteAccount() status = %v, want %v", recorder.Code, testValue.wantStatus)
			}
		})
	}
}
