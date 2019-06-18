package operation

import (
	"errors"
	"hospital/storage"
	"log"
	"os"
	"strconv"
	"time"
)

var ErrTimeout = errors.New("timeout")

func getOperations(surgeonID string) ([]*storage.Operation, error) {
	timeoutTime, err := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT_SECONDS"))
	if err != nil {
		panic(err)
	}

	queryTime, err := strconv.Atoi(os.Getenv("QUERY_INTERVAL_SECONDS"))
	if err != nil {
		panic(err)
	}

	firedOpsStr, err := storage.GetOperation(surgeonID)
	if err != nil {
		return make([]*storage.Operation, 0), err
	}
	if len(firedOpsStr) != 0 {
		return firedOpsStr, nil
	}

	ticker := time.NewTicker(time.Duration(queryTime) * time.Second)
	defer ticker.Stop()

	timeout := time.NewTicker(time.Duration(timeoutTime) * time.Second)
	defer timeout.Stop()

	for {
		select {

		case <-timeout.C:
			log.Printf("Received timeout")
			return make([]*storage.Operation, 0), ErrTimeout

		case <-ticker.C:
			firedOpsStr, err := storage.GetOperation(surgeonID)
			if err != nil {
				return make([]*storage.Operation, 0), err
			}

			if len(firedOpsStr) != 0 {
				return firedOpsStr, nil
			}
		}
	}
}
