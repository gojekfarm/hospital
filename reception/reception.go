package reception

import (
	"log"

	"hospital/doctor"
	"hospital/storage"
)

func alertReceiver(receivedObj AlertReceived) string {
	id := storage.InsertAlertUnique(receivedObj.Alerts[0].Labels.Alertname,
		receivedObj.Alerts[0].StartsAt, receivedObj.Alerts[0].Labels.Job, receivedObj.Alerts[0].Labels.Instance, receivedObj.Alerts[0].Status)

	if receivedObj.Alerts[0].Status == "firing" {
		return callTODoctor(id, receivedObj.Alerts[0].Labels.Alertname, receivedObj.Alerts[0].Labels.Job)
	}
	return "was resolved"
}

func callTODoctor(id int, alertname, applicationID string) string {
	err := doctor.ResolveAlert(id, alertname, applicationID)
	if err != nil {
		log.Println(err)
	}
	return "was firing"
}

// AlertReceived struct exported.
type AlertReceived struct {
	Receiver          string                   `json:"reciever"`
	Status            string                   `json:"status"`
	Alerts            []Alert                  `json:"alerts"`
	GroupLabels       string                   `json:"groupLables"`
	CommonLabels      string                   `json:"commonLables"`
	CommonAnnotations struct{ Summary string } `json:"commonAnnotations"`
	ExternalURL       string                   `json:"externalURL"`
	Version           string                   `json:"version"`
	GroupKey          string                   `json:"groupKey"`
}

// Alert struct exported.
type Alert struct {
	Status       string                   `json:"status"`
	Labels       Label                    `json:"labels"`
	Annotations  struct{ Summary string } `json:"annotations"`
	StartsAt     string                   `json:"startsAT"`
	EndsAt       string                   `json:"endsAT"`
	GeneratorURL string                   `json:"generatorURL"`
}

// Label struct exported.
type Label struct {
	Alertname string `json:"alertname"`
	Backend   string `json:"backend"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Severity  string `json:"severity"`
}
