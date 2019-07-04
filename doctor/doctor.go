package doctor

import (
	"hospital/storage"
)

// ResolveAlert receives alert from reception and generate operations.
func ResolveAlert(alertID int, alertName, applicationID string) error {
	script, err := storage.GetScript(alertName)
	if err != nil {
		return err
	}

	storage.InsertOperation(alertID, applicationID, script, "CRITICAL")
	return nil
}
