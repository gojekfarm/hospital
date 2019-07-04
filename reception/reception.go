package reception

import (
	"log"

	"hospital/doctor"
	"hospital/storage"
)

func receiveAlert(receivedObj AlertReceived) string {
	id := storage.InsertAlertUnique(receivedObj.Message,
		receivedObj.Time, receivedObj.ID, receivedObj.Level)

	if receivedObj.PreviousLevel == "OK" {
		return callTODoctor(id, receivedObj.Message, receivedObj.ID)
	}
	return "was OK"
}

func callTODoctor(id int, alertname, applicationID string) string {
	err := doctor.ResolveAlert(id, alertname, applicationID)
	if err != nil {
		log.Println(err)
	}
	return "was CRITICAL"
}

// AlertReceived struct exported.
type AlertReceived struct {
	ID            string `json:"id"`
	Message       string `json:"message"`
	Details       string `json:"details"`
	Time          string `json:"time"`
	Duration      string `json:"-"`
	Level         string `json:"level"`
	Data          string `json:"-"`
	PreviousLevel string `json:"previousLevel"`
}
