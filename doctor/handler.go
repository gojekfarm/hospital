package doctor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DoctorHandler recieves alerts
func DoctorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		var alertType AlertName
		json.Unmarshal(body, &alertType)
		resp := scriptGenerator(alertType.Alertname)
		fmt.Fprintf(w, resp)
	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}
