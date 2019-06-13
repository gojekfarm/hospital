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
	http.HandleFunc("/v1/reception", reception.Handler)
	http.HandleFunc("/v1/doctor", doctor.Handler)
}
