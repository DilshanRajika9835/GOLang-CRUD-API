package main

import (
	//This is customer route  file
	"BackEnd/CustomerRout"
	//This is import to write values to the console
	"fmt"
	// Go MySQL Driver is aMySQL driver for Go's database/sql package
	_ "github.com/go-sql-driver/mysql"
	// This implements a request router and dispatcher for matching  incoming requests to their
	"github.com/gorilla/mux"
	//This allows us to log errors and other issues
	"log"
	//This gives us HTTP client and server implementation for building the actual API
	"net/http"
)

func main() {
	//Creating a route object
	r := mux.NewRouter()
	fmt.Println("Server Running...")
	// handle all requests
	r.HandleFunc("/api/customer", CustomerRout.SaveCustomer).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/customer", CustomerRout.DeleteCustomer).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/customer", CustomerRout.UpdateCustomer).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/customer", CustomerRout.AllCustomer).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/customer/{id}", CustomerRout.SearchCustomer).Methods("GET", "OPTIONS")

	//Bind to a port and pass our router in
	//This  listen on the TCP network address :8080
	log.Fatal(http.ListenAndServe(":8000", r))

}
