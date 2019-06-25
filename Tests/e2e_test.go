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
	"os"
	"strings"
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

	err = storage.DeleteOps("receptiontest")
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
	applicationID := "surgeontest"
	script := "ls"

	storage.Initialize()
	err := storage.InsertScript("test", script)
	if err != nil {
		panic(err)
	}

	opInsertedID := storage.InsertOperation(-1, applicationID, script, "firing")

	var mux = http.NewServeMux()
	mux.HandleFunc(routes.OperationAPIPath, operation.Handler)
	mux.HandleFunc(routes.ReportAPIPath, report.Handler)

	srv := httptest.NewServer(mux)
	defer srv.Close()

	urls := strings.Split(srv.URL, ":")

	port := os.Getenv("PORT")
	os.Setenv("PORT", urls[2])
	defer os.Setenv("PORT", port)

	oldApplicationID := os.Getenv("APPLICATION_ID")
	os.Setenv("APPLICATION_ID", applicationID)
	defer os.Setenv("APPLICATION_ID", oldApplicationID)

	surgeon.SetVariableFromEnv()

	err = surgeon.MakeRequest()

	if err != nil {
		panic(err)
	}

	status, err := storage.GetOpStatus(opInsertedID)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "completed", status, "completed status is expected")

	err = storage.DeleteOps(applicationID)
	if err != nil {
		panic(err)
	}

}

type receptionResponse struct {
	Status string `json:"status"`
}
