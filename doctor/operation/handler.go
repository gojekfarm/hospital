package operation

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
		var oprequest operationRequest
		err := json.Unmarshal(body, &oprequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
		resp, err := getOperations(oprequest.SurgeonID)
		if err != nil {
			if err == ErrTimeout {
				// Do special handling.
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		fmt.Fprintf(w, resp)

	default:
		fmt.Fprintf(w, "Only post method supported.")
	}
}

type operationRequest struct {
	SurgeonID string `json:"surgeonID"`
}
