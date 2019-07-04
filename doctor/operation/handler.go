package operation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handle receives alerts.
func Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println(err)
			return
		}

		var oprequest operationRequest
		if err = json.Unmarshal(body, &oprequest); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println(err)
			return
		}

		ops, err := getOperations(oprequest.ApplicationID)
		if err != nil {
			if err == ErrNoContent {
				w.WriteHeader(http.StatusNoContent)
				fmt.Fprintln(w, "Polling timeout")
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(ops)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, string(jsonStr))

	default:
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		fmt.Fprintf(w, "Only get method supported.")
	}
}

type operationRequest struct {
	ApplicationID string `json:"applicationID"`
}
