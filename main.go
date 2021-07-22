package main

import (
	"fmt"
	"log"
	"net/http"
	. "simple-rest-api-mux-golang/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// init router
	r := mux.NewRouter()

	r.HandleFunc("/api/books", GetBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", GetBook).Methods("GET")
	r.HandleFunc("/api/book", CreateBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", DeleteBook).Methods("DELETE")

	// server := new(http.Server)
	// server.Addr = ":9000"
	// Init()
	// fmt.Println("Server Running ON ", server.Addr)
	// log.Fatal(server.ListenAndServe())
	Init()
	fmt.Println("Server Running ON http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
