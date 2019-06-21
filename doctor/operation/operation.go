package operation

import (
	"errors"
	"hospital/storage"
	"os"
	"strconv"
	"time"
)

// ErrNoContent for timeout error.
var ErrNoContent = errors.New("timeout")

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
			return make([]*storage.Operation, 0), ErrNoContent

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
