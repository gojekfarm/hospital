package test

import (
	"bytes"
	"encoding/json"
	"hospital/doctor/operation"
	"hospital/doctor/report"
	"hospital/reception"
	"hospital/routes"
	"hospital/storage"
	"hospital/surgeon"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceptionFiring(t *testing.T) {
	storage.Initialize()

	lable := reception.Label{Alertname: "test", Instance: "test", Job: "receptiontest"}
	alert := reception.Alert{Status: "firing", StartsAt: "test", Labels: lable}
	var respAlert = reception.AlertReceived{Alerts: []reception.Alert{alert}}

	jsonStr, err := json.Marshal(&respAlert)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/reception", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(reception.Handler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code, "OK response is expected")

	var resp receptionResponse
	body, _ := ioutil.ReadAll(rr.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		t.Fatal(err)
	}
	expected := "was firing"
	if resp.Status != expected {
		t.Errorf("alert send failed")
	}

	_, err = db.Exec("DELETE FROM operations WHERE application_id = $1", "receptiontest")
	if err != nil {
		panic(err)
	}

}

func TestReceptionResolved(t *testing.T) {
	storage.Initialize()
	lable := reception.Label{Alertname: "test", Instance: "test", Job: "receptiontest"}
	alert := reception.Alert{Status: "resolved", StartsAt: "test", Labels: lable}
	var respAlert = reception.AlertReceived{Alerts: []reception.Alert{alert}}

	jsonStr, err := json.Marshal(&respAlert)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/reception", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(reception.Handler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code, "OK response is expected")

	var resp receptionResponse
	body, _ := ioutil.ReadAll(rr.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		t.Fatal(err)
	}
	expected := "was resolved"
	if resp.Status != expected {
		t.Errorf("alert send failed")
	}
}

func TestSurgeon(t *testing.T) {
	surgeon.ApplicationID = "surgeontest"
	script := "echo \"e2e tests\""

	storage.Initialize()
	db := storage.ReturnDbInstance()
	err := storage.InsertScript("test", script)
	if err != nil {
		panic(err)
	}

	operationID := storage.InsertOperation(-1, surgeon.ApplicationID, script, "firing")

	var mux = http.NewServeMux()
	mux.HandleFunc(routes.OperationAPIPath, operation.Handler)
	mux.HandleFunc(routes.ReportAPIPath, report.Handler)

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	surgeon.HospitalURL = testServer.URL

	err = surgeon.MakeRequest()
	if err != nil {
		panic(err)
	}

	var logs string
	err = db.QueryRow(`SELECT logs FROM operations WHERE id = $1`,
		operationID).Scan(&logs)
	if err != nil {
		panic(err)
	}

	// Checking if desired status is found or not.
	assert.Equal(t, "e2e tests\n", logs, "\"e2e tests\" is expected in logs")

	var status string
	err = db.QueryRow(`SELECT status FROM operations WHERE id = $1`,
		operationID).Scan(&status)
	if err != nil {
		panic(err)
	}

	// Checking if desired status is found or not.
	assert.Equal(t, "completed", status, "completed status is expected")

	_, err = db.Exec("DELETE FROM operations WHERE application_id = $1", surgeon.ApplicationID)
	if err != nil {
		panic(err)
	}

}

type receptionResponse struct {
	Status string `json:"status"`
}
