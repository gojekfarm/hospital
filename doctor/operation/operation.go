package operation

import (
	"fmt"
	"hospital/storage"
)

func getOperation(surgeonID string) string {
	firedOperations := storage.GetOperation(surgeonID)
	fmt.Println(firedOperations)
	return firedOperations
}
