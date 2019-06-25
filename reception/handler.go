package reception

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler recieves alerts
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)

		var receivedObject AlertReceived
		err := json.Unmarshal(body, &receivedObject)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		resp := alertReceiver(receivedObject)

		response := struct {
			Status string `json:"status"`
		}{
			Status: resp,
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintf(w, string(responseJSON))
	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}
