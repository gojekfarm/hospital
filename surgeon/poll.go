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

var errServer = errors.New("Server error")

func MakeRequest() error {
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

		var ops []Operation
		err := json.Unmarshal(body, &ops)
		if err != nil {
			log.Println(err)
		}

		RunScripts(ops)
	} else if resp.StatusCode == 204 {
		log.Println("No operations to execute.")
	} else {
		return errServer
	}
	return nil
}

// RunScripts runs all operations in script.
func RunScripts(ops []Operation) {
	for _, op := range ops {
		cmd := exec.Command("sh", "-c", op.Script)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		outStr, errStr := stdout.String(), stderr.String()

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

type Operation struct {
	ID     int    `json:"id"`
	Script string `json:"script"`
}
