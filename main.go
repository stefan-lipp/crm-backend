package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Company   string `json:"company,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

var customers = []Customer{
	{Id: 1, Name: "Tim Cook", Role: "CEO", Company: "Apple", Email: "tim.cook@apple.com", Phone: "+1 123 456 789", Contacted: true},
	{Id: 2, Name: "Larry Page", Role: "Founder", Company: "Google", Email: "larry.page@google.com", Phone: "+1 234 567 890", Contacted: true},
	{Id: 3, Name: "Elon Musk", Role: "Founder & CEO", Company: "Tesla", Email: "elon.musk@tesla.com", Phone: "+1 345 678 901", Contacted: false},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	for _, customer := range customers {
		if customer.Id == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add customer...")
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update customer...")
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete customer...")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting the CRM backend server on port 3000...")
	http.ListenAndServe(":3000", router)
}
