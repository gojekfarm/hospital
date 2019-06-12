package main

import (
	"fmt"
	"hospital/doctor"
	"hospital/reception"
	"net/http"
)

// PingHandler gives response pong.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

//Startserver handles our whole routing and server
func Startserver(port string) {
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/reception", reception.ReceptionHandler)
	http.HandleFunc("/doctor", doctor.DoctorHandler)
	// http.HandleFunc("/operation", OperationHandler)
	// http.HandleFunc("/reporting", ReportingHandler)
	http.ListenAndServe(":"+port, nil)
}
