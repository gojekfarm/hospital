package operation

import (
	"errors"
	"fmt"
	"hospital/storage"
	"log"
	"time"
)

var ErrTimeout = errors.New("timeout")

func getOperations(surgeonID string) (string, error) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	timeout := time.NewTicker(30 * time.Second)
	defer timeout.Stop()

	for {
		select {
		// case <-ctx.Done():
		// 	log.Printf("Received context cancel")
		// 	return ""
		case <-timeout.C:
			log.Printf("Received timeout")
			return "", ErrTimeout

		case <-ticker.C:
			firedOpsStr, err := storage.GetOperation(surgeonID)
			if err != nil {
				return "", err
			}
			fmt.Println("ticking")
			if firedOpsStr != "[]" {
				return firedOpsStr, nil
			}
		}
	}
}
