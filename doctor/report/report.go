package report

import (
	"bytes"
	"encoding/json"
	"hospital/storage"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func reportStatus(reportReq reportRequest) error {
	alertname, err1 := storage.AlertNameFromOpID(reportReq.ID)
	if err1 != nil {
		log.Println(err1)
	}

	applicationID, err2 := storage.GetApplicationID(reportReq.ID)
	if err2 != nil {
		log.Println(err2)
	}

	if err1 == nil && err2 == nil {
		slackReport(applicationID, alertname, reportReq.Status, reportReq.Logs)
	}

	err := storage.RecordStatus(reportReq.ID, reportReq.Status, reportReq.Logs)
	return err
}

func slackReport(applicationID, alertname, status, logs string) {
	webhookURL := os.Getenv("SLACK_URL")

	attachment1 := attachment{}

	color := "good"
	if status == "failed" {
		color = "danger"
	}
	attachment1.Color = &color

	attachment1.addField(field{Title: "Host Name",
		Value: applicationID,
		Short: true})

	attachment1.addField(field{Title: "Alert Name",
		Value: alertname,
		Short: true})

	// if size := len(logs); size > 50 {
	// 	logs = logs[size-50:]
	// }

	attachment1.addField(field{Title: "Logs",
		Value: "`" + "`" + "`" + logs + "`" + "`" + "`",
		Short: false})

	payloadSent := payload{
		Text:        "Reporting the logs of command executed:",
		Attachments: []attachment{attachment1},
	}

	jsonStr, err := json.Marshal(payloadSent)
	if err != nil {
		log.Println(err)
		return
	}

	// Making request to Slack.
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
