package surgeon

import (
	"bytes"
	"encoding/json"
	"hospital/routes"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func makeRequest() {
	var jsonStr = []byte(`{"surgeonID":"` + surgeonID + `"}`)
	req, err := http.NewRequest("GET", url+routes.OperationAPIPath, bytes.NewBuffer(jsonStr))
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

	if resp.StatusCode == 200 {
		var ops []operation
		err := json.Unmarshal(body, &ops)
		if err != nil {
			log.Println(err)
		}

		runScripts(ops)
	} else {
		log.Println("No operations to execute.")
	}
}

func runScripts(ops []operation) {
	for _, op := range ops {
		out, err := exec.Command("sh", "-c", op.Script).Output()
		exitCode := 0
		if err != nil {
			log.Println(err)
			status, _ := err.(*exec.ExitError)
			exitCode = status.ExitCode()
		}

		status := "completed"
		if exitCode != 0 {
			status = "failed"
		}
		log.Println(string(out))

		makeReport(op.ID, status, string(out))
	}
}

type operation struct {
	ID     int    `json:"id"`
	Script string `json:"script"`
}

func makeReport(id int, status, logs string) {
	reqBody := reportReq{id, status, logs}

	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", url+routes.ReportAPIPath, bytes.NewBuffer(jsonStr))
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
