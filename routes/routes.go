package routes

import (
	"hospital/dashboard"
	"hospital/doctor/operation"
	"hospital/doctor/report"
	"hospital/healthcheck"

	"hospital/reception"

	"github.com/gorilla/mux"
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
func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(PingAPIPath, healthcheck.Handle)
	router.HandleFunc(ReceptionAPIPath, reception.Handle)
	router.HandleFunc(ReportAPIPath, report.Handle)
	router.HandleFunc(OperationAPIPath, operation.Handle)

	router.HandleFunc("/dashboard", dashboard.HandleDashboard)
	router.HandleFunc("/dashboard/logs", dashboard.HandleLogs)
	router.HandleFunc("/dashboard/logs/{id}", dashboard.HandleOneLog)
	router.HandleFunc("/dashboard/insert", dashboard.HandleInsert)
	router.HandleFunc("/dashboard/remove/", dashboard.HandleRemove)
	router.HandleFunc("/dashboard/summary", dashboard.HandleSummary)
	router.HandleFunc("/dashboard/summary/{id}", dashboard.HandleOneSummary)

	return router
}
