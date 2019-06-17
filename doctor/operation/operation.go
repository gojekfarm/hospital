package operation

import (
	"hospital/storage"
)

func getOperation(surgeonID string) string {
	firedOperations := storage.GetOperation(surgeonID)
	return firedOperations
}
