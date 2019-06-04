package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Only post methods supported.")
	case "POST":
		cmdStr := "thumbnail.sh"
		cmd := exec.Command("/bin/sh", cmdStr)
		_, err := cmd.Output()

		if err != nil {
			println(err.Error())
			return
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
