package test

import (
	"bytes"
	"encoding/json"
	"hospital/reception"
	"hospital/storage"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceptionFiring(t *testing.T) {
	storage.Initialize()
	var jsonStr = []byte(`{"receiver":"pepper","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"test","backend":"localnodes","instance":"test","job":"haproxy","severity":"page"},"annotations":{"summary":"Current queue is greater than 100"},"startsAt":"test","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://Dilips-MacBook-Pro.local:9090/graph?g0.expr=haproxy_backend_current_queue%7Bbackend%3D%22localnodes%22%2Cinstance%3D%22localhost%3A9101%22%2Cjob%3D%22haproxy%22%7D+%3E+100\u0026g0.tab=1"}],"groupLabels":{},"commonLabels":{"alertname":"queue_exceed","backend":"localnodes","instance":"localhost:9101","job":"haproxy","severity":"page"},"commonAnnotations":{"summary":"Current queue is greater than 100"},"externalURL":"http://Dilips-MacBook-Pro.local:9093","version":"4","groupKey":"{}:{}"}
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
}

func TestReceptionResolved(t *testing.T) {
	storage.Initialize()
	var jsonStr = []byte(`{"receiver":"pepper","status":"firing","alerts":[{"status":"resolved","labels":{"alertname":"test","backend":"localnodes","instance":"test","job":"haproxytest","severity":"page"},"annotations":{"summary":"Current queue is greater than 100"},"startsAt":"test","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://Dilips-MacBook-Pro.local:9090/graph?g0.expr=haproxy_backend_current_queue%7Bbackend%3D%22localnodes%22%2Cinstance%3D%22localhost%3A9101%22%2Cjob%3D%22haproxy%22%7D+%3E+100\u0026g0.tab=1"}],"groupLabels":{},"commonLabels":{"alertname":"queue_exceed","backend":"localnodes","instance":"localhost:9101","job":"haproxy","severity":"page"},"commonAnnotations":{"summary":"Current queue is greater than 100"},"externalURL":"http://Dilips-MacBook-Pro.local:9093","version":"4","groupKey":"{}:{}"}
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
	storage.Initialize()
	err := storage.InsertScript("test", "ls")
	if err != nil {
		panic(err)
	}
}

type receptionResponse struct {
	Status string `json:"status"`
}
