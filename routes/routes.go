package routes

import (
	"hospital/doctor/operation"
	"hospital/healthCheck"
	"hospital/reception"
	"net/http"
	"time"
)

//Routes handles our whole routing and server
func Routes() {
	http.HandleFunc("/ping", healthCheck.Handler)
	http.HandleFunc("/v1/reception", reception.Handler)
	http.Handle("/v1/operation",
		http.TimeoutHandler(http.HandlerFunc(operation.Handler), 30*time.Second,
			"Your request has timed out. :("))
}
