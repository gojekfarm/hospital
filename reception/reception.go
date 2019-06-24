package reception

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"

	"hospital/doctor"
	"hospital/storage"
)

func alertReceiver(receivedObj received) string {
	id := storage.InsertAlertUnique(receivedObj.Alerts[0].Labels.Alertname,
		receivedObj.Alerts[0].StartsAt, receivedObj.Alerts[0].Labels.Job, receivedObj.Alerts[0].Labels.Instance, receivedObj.Alerts[0].Status)

	if receivedObj.Alerts[0].Status == "firing" {
		return callTODoctor(id, receivedObj.Alerts[0].Labels.Alertname, receivedObj.Alerts[0].Labels.Job)
	}
	return `{"status" : "was resolved"}`
}

func callTODoctor(id int, alertname, applicationID string) string {
	var jsonStr = []byte(`{"id":` + strconv.Itoa(id) + `, "alertname":"` + alertname + `", "applicationID":"` + applicationID + `"}`)
	req, err := http.NewRequest("POST", "/doctor", bytes.NewBuffer(jsonStr))
	if err != nil {
		return `{"status" : "reception failed"}`
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(doctor.Handler)

	handler.ServeHTTP(rr, req)
	log.Println(rr.Body)
	return `{"status" : "was firing"}`
}

type received struct {
	Receiver          string                   `json:"reciever"`
	Status            string                   `json:"status"`
	Alerts            []alert                  `json:"alerts"`
	GroupLabels       string                   `json:"groupLables"`
	CommonLabels      string                   `json:"commonLables"`
	CommonAnnotations struct{ Summary string } `json:"commonAnnotations"`
	ExternalURL       string                   `json:"externalURL"`
	Version           string                   `json:"version"`
	GroupKey          string                   `json:"groupKey"`
}

type alert struct {
	Status       string                   `json:"status"`
	Labels       label                    `json:"labels"`
	Annotations  struct{ Summary string } `json:"annotations"`
	StartsAt     string                   `json:"startsAT"`
	EndsAt       string                   `json:"endsAT"`
	GeneratorURL string                   `json:"generatorURL"`
}

type label struct {
	Alertname string `json:"alertname"`
	Backend   string `json:"backend"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Severity  string `json:"severity"`
}
