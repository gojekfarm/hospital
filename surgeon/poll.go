package surgeon

import (
	"bytes"
	"encoding/json"
	"errors"
	"hospital/routes"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

var errServer = errors.New("server error")

func makeRequest() error {
	var jsonStr = []byte(`{"applicationID":"` + applicationID + `"}`)
	req, err := http.NewRequest("GET", url+routes.OperationAPIPath, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return errServer
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
	} else if resp.StatusCode == 204 {
		log.Println("No operations to execute.")
	} else {
		return errServer
	}
	return nil
}

func runScripts(ops []operation) {
	for _, op := range ops {
		cmd := exec.Command("sh", "-c", op.Script)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

		exitCode := 0
		if err != nil {
			log.Println(err)
			status, _ := err.(*exec.ExitError)
			exitCode = status.ExitCode()
		}

		logs := outStr
		status := "completed"
		if exitCode != 0 {
			logs = errStr
			status = "failed"
		}
		log.Println(logs)

		makeReport(op.ID, status, logs)
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
