package doctor

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
		var alertDetails alert
		err := json.Unmarshal(body, &alertDetails)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		resp := resolveAlert(alertDetails.ID, alertDetails.Alertname, alertDetails.ApplicationID)
		respStr := `{"status" : "` + resp + `"}`
		fmt.Fprintf(w, respStr)
	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}

type alert struct {
	ID            int    `json:"id"`
	Alertname     string `json:"alertname"`
	ApplicationID string `json:"applicationID"`
}
