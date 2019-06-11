package reception

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"../dbprovider"
	"../doctor"
)

// ReceptionHandler recieves alerts
func ReceptionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		var receivedObject received
		json.Unmarshal(body, &receivedObject)
		resp := alertReceiver(receivedObject)
		fmt.Fprintf(w, resp)
	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}

func alertReceiver(receivedObj received) string {
	dbprovider.InsertAlertUnique(receivedObj.Alerts[0].Labels.Alertname, receivedObj.Alerts[0].StartsAt, receivedObj.Alerts[0].Labels.Instance, receivedObj.Alerts[0].Status)

	if receivedObj.Alerts[0].Status == "firing" {
		return callTODoctor(receivedObj.Alerts[0].Labels.Alertname)
	}
	return "success"
}

func callTODoctor(alertname string) string {
	var jsonStr = []byte(`{"alertname":"` + alertname + `"}`)
	req, err := http.NewRequest("POST", "/doctor", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err.Error()
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(doctor.DoctorHandler)

	handler.ServeHTTP(rr, req)
	return "success"
}

type received struct {
	Receiver          string
	Status            string
	Alerts            []alert
	GroupLabels       string
	CommonLabels      string
	CommonAnnotations string
	ExternalURL       string
	Version           string
	GroupKey          string
}

type alert struct {
	Status       string
	Labels       label
	Annotations  string
	StartsAt     string
	EndsAt       string
	GeneratorURL string
}

type label struct {
	Alertname string
	Backend   string
	Instance  string
	Job       string
	Severity  string
}
