package main

import (
	"net/http"

	"./reception"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/reception", reception.ReceptionHandler)
	// r.HandleFunc("/operation", OperationHandler)
	// r.HandleFunc("/reporting", ReportingHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8088", nil)
}
