package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertReceiver(t *testing.T) {
	var jsonStr = []byte(`{"receiver":"pepper","status":"firing","alerts":[{"status":"test","labels":{"alertname":"test","backend":"localnodes","instance":"test","job":"haproxy","severity":"page"},"annotations":{"summary":"Current queue is greater than 100"},"startsAt":"test","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://Dilips-MacBook-Pro.local:9090/graph?g0.expr=haproxy_backend_current_queue%7Bbackend%3D%22localnodes%22%2Cinstance%3D%22localhost%3A9101%22%2Cjob%3D%22haproxy%22%7D+%3E+100\u0026g0.tab=1"}],"groupLabels":{},"commonLabels":{"alertname":"queue_exceed","backend":"localnodes","instance":"localhost:9101","job":"haproxy","severity":"page"},"commonAnnotations":{"summary":"Current queue is greater than 100"},"externalURL":"http://Dilips-MacBook-Pro.local:9093","version":"4","groupKey":"{}:{}"}
	`)
	req, err := http.NewRequest("POST", "/accept", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(acceptHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code, "OK response is expected")

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "success"
	if rr.Body.String() != expected {
		t.Errorf("alert send failed")
	}
}

func TestDatabase(t *testing.T) {
	db, err := sql.Open("mysql",
		"root:toor@tcp(127.0.0.1:3306)/secretory")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var jsonStr = []byte(`{"receiver":"pepper","status":"firing","alerts":[{"status":"test","labels":{"alertname":"test","backend":"localnodes","instance":"test","job":"haproxy","severity":"page"},"annotations":{"summary":"Current queue is greater than 100"},"startsAt":"test","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://Dilips-MacBook-Pro.local:9090/graph?g0.expr=haproxy_backend_current_queue%7Bbackend%3D%22localnodes%22%2Cinstance%3D%22localhost%3A9101%22%2Cjob%3D%22haproxy%22%7D+%3E+100\u0026g0.tab=1"}],"groupLabels":{},"commonLabels":{"alertname":"queue_exceed","backend":"localnodes","instance":"localhost:9101","job":"haproxy","severity":"page"},"commonAnnotations":{"summary":"Current queue is greater than 100"},"externalURL":"http://Dilips-MacBook-Pro.local:9093","version":"4","groupKey":"{}:{}"}
	`)
	var receivedObject received
	json.Unmarshal(jsonStr, &receivedObject)

	insert, err := db.Query("INSERT INTO alerts(alertname, startsAT, address, status) VALUES ( ? , ? , ? ,?)",
		receivedObject.Alerts[0].Labels.Alertname, receivedObject.Alerts[0].StartsAt, receivedObject.Alerts[0].Labels.Instance, receivedObject.Alerts[0].Status)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	var id int
	err = db.QueryRow("select id from alerts where alertname = ? and startsAT = ? and address = ?",
		receivedObject.Alerts[0].Labels.Alertname, receivedObject.Alerts[0].StartsAt, receivedObject.Alerts[0].Labels.Instance).Scan(&id)
	if err != nil {
		t.Errorf("database test failed!")
	}

}
