package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/emp", GetEmployees).Methods("GET")
	r.HandleFunc("/emp/{eid}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/emp", CrateEmployee).Methods("POST")
	r.HandleFunc("/emp/{eid}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/emp/{eid}", DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
