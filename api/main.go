package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerV1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n v1"))
}

func HandlerV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n v2"))
}
func ServerV1() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1", HandlerV1)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
func ServerV2() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v2", HandlerV2)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":9000", r))
}
func main() {
	go ServerV1()
	ServerV2()
}
