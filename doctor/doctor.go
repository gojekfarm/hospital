package doctor

import (
	"hospital/storage"
)

func resolveAlert(alertID int, alertName, jobName string) string {
	script := storage.GetScript(alertName)
	if script != "no script" {
		storage.InsertOperation(alertID, jobName, script, "firing")
	}
	return "done"
}
