package routes

import (
	"hospital/doctor/operation"
	"hospital/healthCheck"
	"hospital/reception"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Routes handles our whole routing and server
func Routes() {
	http.HandleFunc("/ping", healthCheck.Handler)
	http.HandleFunc("/v1/reception", reception.Handler)

	timeoutTime, err := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	http.Handle("/v1/operation",
		http.TimeoutHandler(http.HandlerFunc(operation.Handler), time.Duration(timeoutTime)*time.Second,
			"Your request has timed out. :("))
}
