package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

func longPolling() {
	ticker := time.NewTicker(35 * time.Second)
	defer ticker.Stop()

	makeRequest()

	for {
		select {

		case <-ticker.C:
			makeRequest()
		}
	}

}

func makeRequest() {
	var jsonStr = []byte(`{"surgeonID":"` + surgeonID + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if resp.StatusCode == 200 {
		var ops []operation
		err := json.Unmarshal(body, &ops)
		if err != nil {
			panic(err)
		}

		runScripts(ops)
	}
}

func runScripts(ops []operation) {
	for _, op := range ops {
		out, err := exec.Command("sh", "-c", op.Script).Output()
		if err != nil {
			panic(err)
		}
		println(string(out))
	}
}

type operation struct {
	ID     int    `json:"id"`
	Script string `json:"script"`
}
