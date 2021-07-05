package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID        string `json:"id"`
	Isbn      string `json:"isbn"`
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
}

var employees []Employee

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

func main() {

	employees = append(employees, Employee{ID: "1", Isbn: "12345", Firstname: "Anand", Lastname: "Pandey"})
	employees = append(employees, Employee{ID: "2", Isbn: "13245", Firstname: "Siddharth", Lastname: "Soni"})
	r := mux.NewRouter()
	r.HandleFunc("/employees/{id}", getEmployee).Methods("GET")

	fmt.Println("Server has started on 8081: ")
	log.Fatal(http.ListenAndServe(":8081", r))
}
