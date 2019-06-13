package routes

import (
	"hospital/doctor"
	"hospital/healthCheck"
	"hospital/reception"
	"net/http"
)


//Routes handles our whole routing and server
func Routes() {
	http.HandleFunc("/ping", healthCheck.Handler)
	http.HandleFunc("/reception", reception.Handler)
	http.HandleFunc("/doctor", doctor.Handler)
}
