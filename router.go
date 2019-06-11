package main

import (
	"net/http"

	"hospital/doctor"
	"hospital/reception"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/reception", reception.ReceptionHandler)
	r.HandleFunc("/doctor", doctor.DoctorHandler)
	// r.HandleFunc("/operation", OperationHandler)
	// r.HandleFunc("/reporting", ReportingHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8088", nil)
}
