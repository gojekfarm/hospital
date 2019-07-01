package routes

import (
	"hospital/dashboard"
	"hospital/doctor/operation"
	"hospital/doctor/report"
	"hospital/healthcheck"

	"hospital/reception"
	"net/http"
)

var (
	// OperationAPIPath for getting operation for surgeon.
	OperationAPIPath = "/v1/operation"
	// PingAPIPath for health check.
	PingAPIPath = "/ping"
	// ReceptionAPIPath for receiving alert via alertmanger.
	ReceptionAPIPath = "/v1/reception"
	// ReportAPIPath used by surgeon for reporting results to doctor.
	ReportAPIPath = "/v1/report"
)

//Routes handles our whole routing and server
func Routes() {
	http.HandleFunc(PingAPIPath, healthcheck.Handle)
	http.HandleFunc(ReceptionAPIPath, reception.Handle)
	http.HandleFunc(ReportAPIPath, report.Handle)
	http.HandleFunc(OperationAPIPath, operation.Handle)

	http.HandleFunc("/dashboard", dashboard.HandleDashboard)
	http.HandleFunc("/dashboard/logs", dashboard.HandleLogs)
	http.HandleFunc("/dashboard/insert", dashboard.HandleInsert)
	http.HandleFunc("/dashboard/remove/", dashboard.HandleRemove)
}
