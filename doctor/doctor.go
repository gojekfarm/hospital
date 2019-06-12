package doctor

import (
	"hospital/dbprovider"
)

func scriptGenerator(alertType string) string {
	script := dbprovider.GetScript(alertType)
	return script
}

type AlertName struct {
	Alertname string `json: "alertname"`
}
