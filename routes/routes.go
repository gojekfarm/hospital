package routes

import (
	"hospital/doctor/operation"
	"hospital/doctor/report"
	"hospital/healthCheck"

	"hospital/reception"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	OperationAPIPath = "/v1/operation"
	PingAPIPath      = "/ping"
	ReceptionAPIPath = "/v1/reception"
	ReportAPIPath    = "/v1/report"
)

//Routes handles our whole routing and server
func Routes() {
	http.HandleFunc(PingAPIPath, healthCheck.Handler)
	http.HandleFunc(ReceptionAPIPath, reception.Handler)
	http.HandleFunc(ReportAPIPath, report.Handler)

	timeoutTime, err := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT_SECONDS"))
	if err != nil {
		panic(err)
	}

	http.Handle(OperationAPIPath,
		http.TimeoutHandler(http.HandlerFunc(operation.Handler), time.Duration(timeoutTime)*time.Second,
			"Your request has timed out. :("))
}
