package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to first route </h1>"))

}
//export GO111MODULE="on" 
//go mod tidy 
//go list
//go list -m all
//go list -m -versions github.com/gorilla/mux
//go mod why github/gorilla/mux
//go mod graph
//go mod edit -go 1.16
//go mod vendor
//go run -mod=vendor main.go

