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
<<<<<<< HEAD

	lable := reception.Label{Alertname: "test", Instance: "test", Job: "receptiontest"}
	alert := reception.Alert{Status: "firing", StartsAt: "test", Labels: lable}
	var respAlert = reception.AlertReceived{Alerts: []reception.Alert{alert}}

	jsonStr, err := json.Marshal(&respAlert)
	if err != nil {
		panic(err)
	}

=======
	db := storage.ReturnDbInstance()

	var jsonStr = []byte(`{"receiver":"pepper","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"test","backend":"localnodes","instance":"test","job":"receptiontest","severity":"page"},"annotations":{"summary":"Current queue is greater than 100"},"startsAt":"test","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://Dilips-MacBook-Pro.local:9090/graph?g0.expr=haproxy_backend_current_queue%7Bbackend%3D%22localnodes%22%2Cinstance%3D%22localhost%3A9101%22%2Cjob%3D%22haproxy%22%7D+%3E+100\u0026g0.tab=1"}],"groupLabels":{},"commonLabels":{"alertname":"queue_exceed","backend":"localnodes","instance":"localhost:9101","job":"haproxy","severity":"page"},"commonAnnotations":{"summary":"Current queue is greater than 100"},"externalURL":"http://Dilips-MacBook-Pro.local:9093","version":"4","groupKey":"{}:{}"}
	`)
>>>>>>> 9889765a980c77d6c5b8c5642d7e03a7e41e26c2
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
