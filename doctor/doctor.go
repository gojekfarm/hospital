package doctor

import (
	"hospital/storage"
)

func resolveAlert(alertID int, alertName, applicationID string) string {
	script, err := storage.GetScript(alertName)
	if err != nil {
		return "script not found"
	}

	storage.InsertOperation(alertID, applicationID, script, "firing")

	return "script found"
}
