package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Only post methods supported.")
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		var respString AlertName
		json.Unmarshal(body, &respString)
		log.Println(respString.Alertname + " called!")

		var cmdStr string
		switch respString.Alertname {
		case "queue_exceed":
			cmdStr = "thumbnail.sh"
		default:
			cmdStr = "none"
		}
		if cmdStr != "none" {
			cmd := exec.Command("/bin/sh", cmdStr)
			_, err := cmd.Output()

			if err != nil {
				println(err.Error())
				return
			}
			fmt.Fprintf(w, "success")
		} else {
			fmt.Fprintf(w, "failed")
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func main() {
	http.HandleFunc("/script", scriptHandler)

	log.Println("Go!")
	http.ListenAndServe(":8089", nil)
}

type AlertName struct {
	Alertname string
}
