package routes

import (
	"hospital/healthCheck"
	"hospital/reception"
	"net/http"
)

//Routes handles our whole routing and server
func Routes() {
	http.HandleFunc("/ping", healthCheck.Handler)
	http.HandleFunc("/v1/reception", reception.Handler)
}
