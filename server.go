package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"-"`
}

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Balaji", "Cloud", 10002},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmp.ID = len(employees) + 1

	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
}

func loggingMiddleware(next http.Handler) http.Handler {
	router := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		// if req.Header // if authenticated!
		next.ServeHTTP(w, req)
		// else
		// w.WriteHeader(http.StatusUnauthorized)

		log.Printf("%s %s took %s\n", req.Method, req.URL, time.Since(begin))
	}

	return http.HandlerFunc(router)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")
	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")

	log.Println("Starting server on port: 8000...")
	// http.ListenAndServe(":8000", loggingMiddleware(r))
	http.ListenAndServe(":8000", handlers.CombinedLoggingHandler(os.Stdout, r))
}
