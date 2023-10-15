package main

// TODO
// 1. Create an additional endpoint that updates customer values in a batch (i.e., rather than for a single customer).
// 2. Upgrade the mock database to a real database (e.g., PostgreSQL).
// 3. Deploy the API to the web.

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type customerDetail struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customerDict = map[string]customerDetail{
	"123": {
		"John",
		"Manager",
		"john@gmail.com",
		"1234567890",
		false,
	},
	"456": {
		"Jane",
		"Manager",
		"jane@gmail.com",
		"9991234567890",
		false,
	},
	"789": {
		"Kayn",
		"Director",
		"Kayn@gmail.com",
		"8881234567890",
		true,
	},
}

// The function serves the index.html file to the client.
func showHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

// The function `getCustomerDict` returns a JSON representation of the `customerDict`
// variable.
func getCustomerDict(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerDict)
}

// The function retrieves a single customer's details based on their ID and returns it in
// JSON format.
func getSignleCustomerDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := customerDict[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customerDict[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// The function creates a new customer with a unique ID and adds it to a dictionary if it
// doesn't already exist.
func createSingleCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := uuid.New().String()

	// Parse the request body into a CustomerDetail struct
	var newCustomer customerDetail
	reqBody, err := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Check if the customer already exists
	if _, ok := customerDict[id]; ok {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "Customer already exists"})
		return
	}

	// Add the new customer to the customerDict map
	customerDict[id] = newCustomer

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customerDict)
}

// The function updates a single customer in a dictionary based on the provided ID.
func updateSingleCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer customerDetail
	reqBody, err := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	id := mux.Vars(r)["id"]
	if _, ok := customerDict[id]; ok {
		w.WriteHeader(http.StatusOK)
		customerDict[id] = newCustomer
		json.NewEncoder(w).Encode(customerDict)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// The function deletes a single customer from a dictionary and returns the updated
// dictionary as a JSON response.
func deleteSingleCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if _, ok := customerDict[id]; ok {
		w.WriteHeader(http.StatusOK)
		delete(customerDict, id)
		json.NewEncoder(w).Encode(customerDict)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", showHome).Methods("GET")
	router.HandleFunc("/customers", getCustomerDict).Methods("GET")
	router.HandleFunc("/customers/{id}", getSignleCustomerDetail).Methods("GET")
	router.HandleFunc("/customers", createSingleCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateSingleCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteSingleCustomer).Methods("DELETE")
	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
