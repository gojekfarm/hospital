package surgeon

import (
	"bytes"
	"encoding/json"
	"hospital/routes"
	"io/ioutil"
	"log"
	"net/http"
)

func makeReport(id int, status, logs string) {
	reqBody := reportReq{id, status, logs}

	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", HospitalURL+routes.ReportAPIPath, bytes.NewBuffer(jsonStr))
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

type reportReq struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Logs   string `json:"logs"`
}
