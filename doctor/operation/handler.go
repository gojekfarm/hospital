package operation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler recieves alerts.
func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		body, _ := ioutil.ReadAll(r.Body)

		var oprequest operationRequest

		err := json.Unmarshal(body, &oprequest)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println(err)
			return
		}

		ops, err := getOperations(oprequest.SurgeonID)
		if err != nil {
			if err == ErrTimeout {
				http.Error(w, http.StatusText(http.StatusRequestTimeout), http.StatusRequestTimeout)
				log.Println(err)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(ops)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		fmt.Fprintf(w, string(jsonStr))

	default:
		fmt.Fprintf(w, "Only post method supported.")
	}
}

type operationRequest struct {
	SurgeonID string `json:"surgeonID"`
}
