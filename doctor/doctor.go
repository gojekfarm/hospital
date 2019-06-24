package doctor

import (
	"hospital/storage"
)

func resolveAlert(alertID int, alertName, applicationID string) string {
	script := storage.GetScript(alertName)

	if script != "no script" {
		storage.InsertOperation(alertID, applicationID, script, "firing")
		return "script found"
	}

	return "script not found"
}
