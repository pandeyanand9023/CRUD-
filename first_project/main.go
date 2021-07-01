package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID        string `json:"id"`
	Isbn      string `json:"isbn"`
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
}

var employees []Employee

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range employees {

		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range employees {
		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return

		}
	}
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = strconv.Itoa(rand.Intn(100000000))
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			var employee Employee
			_ = json.NewDecoder(r.Body).Decode(&employee)

			employee.ID = params["id"]
			employees = append(employees, employee)
			json.NewEncoder(w).Encode(employee)

			return
		}
	}
}

func main() {

	employees = append(employees, Employee{ID: "1", Isbn: "12345", Firstname: "Anand", Lastname: "Pandey"})
	employees = append(employees, Employee{ID: "2", Isbn: "13245", Firstname: "Siddharth", Lastname: "Soni"})
	r := mux.NewRouter()
	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/employees", createEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")

	fmt.Println("Server has started on 8080: ")
	log.Fatal(http.ListenAndServe(":8080", r))
}
