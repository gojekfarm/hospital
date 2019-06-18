package surgeon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func makeRequest() {
	var jsonStr = []byte(`{"surgeonID":"` + surgeonID + `"}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
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
		if err != nil {
			log.Println(err)
		}
		println(string(out))
	}
}

type operation struct {
	ID     int    `json:"id"`
	Script string `json:"script"`
}
