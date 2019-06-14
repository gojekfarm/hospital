package operation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handler recieves alerts
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		var oprequest operationRequest
		json.Unmarshal(body, &oprequest)
		resp := getOperation(oprequest.SurgeonID)
		if resp != "null" {
			fmt.Fprintf(w, resp)
		} else {
			respScript := `[]`
			fmt.Fprintf(w, respScript)
		}

	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}

type operationRequest struct {
	SurgeonID string `json:"surgeonID"`
}
