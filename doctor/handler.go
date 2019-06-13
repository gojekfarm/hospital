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
		var alertType alertName
		json.Unmarshal(body, &alertType)
		resp := scriptGenerator(alertType.Alertname)
		respScript := `{"script" : "`+resp+`"}`
		fmt.Fprintf(w, respScript)
	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}

type alertName struct {
	Alertname string `json: "alertname"`
}
