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
	var jsonStr = []byte(`{"receiver":"pepper","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"test","backend":"localnodes","instance":"test","job":"receptiontest","severity":"page"},"annotations":{"summary":"Current queue is greater than 100"},"startsAt":"test","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://Dilips-MacBook-Pro.local:9090/graph?g0.expr=haproxy_backend_current_queue%7Bbackend%3D%22localnodes%22%2Cinstance%3D%22localhost%3A9101%22%2Cjob%3D%22haproxy%22%7D+%3E+100\u0026g0.tab=1"}],"groupLabels":{},"commonLabels":{"alertname":"queue_exceed","backend":"localnodes","instance":"localhost:9101","job":"haproxy","severity":"page"},"commonAnnotations":{"summary":"Current queue is greater than 100"},"externalURL":"http://Dilips-MacBook-Pro.local:9093","version":"4","groupKey":"{}:{}"}
	`)
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
	var jsonStr = []byte(`{"receiver":"pepper","status":"firing","alerts":[{"status":"resolved","labels":{"alertname":"test","backend":"localnodes","instance":"test","job":"receptiontest","severity":"page"},"annotations":{"summary":"Current queue is greater than 100"},"startsAt":"test","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://Dilips-MacBook-Pro.local:9090/graph?g0.expr=haproxy_backend_current_queue%7Bbackend%3D%22localnodes%22%2Cinstance%3D%22localhost%3A9101%22%2Cjob%3D%22haproxy%22%7D+%3E+100\u0026g0.tab=1"}],"groupLabels":{},"commonLabels":{"alertname":"queue_exceed","backend":"localnodes","instance":"localhost:9101","job":"haproxy","severity":"page"},"commonAnnotations":{"summary":"Current queue is greater than 100"},"externalURL":"http://Dilips-MacBook-Pro.local:9093","version":"4","groupKey":"{}:{}"}
	`)
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
