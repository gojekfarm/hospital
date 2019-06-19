package report

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		body, _ := ioutil.ReadAll(r.Body)

		var reportReq reportRequest

		err := json.Unmarshal(body, &reportReq)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = statusReporter(reportReq)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		response := struct {
			Status string `json:"status"`
		}{
			Status: "success",
		}

		responseJson, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintf(w, string(responseJson))

	default:
		fmt.Fprintf(w, "Only post method supported.")
		// error 4xx
	}
}

type reportRequest struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Log    string `json:"log"`
}
