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
		json.Unmarshal(body, &alertDetails)
		resp := resolveAlert(alertDetails.ID, alertDetails.Alertname, alertDetails.JobName)
		respScript := `{"script" : "` + resp + `"}`
		fmt.Fprintf(w, respScript)
	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}

type alert struct {
	ID        int    `json:"id"`
	Alertname string `json:"alertname"`
	JobName   string `json:"jobname"`
}
