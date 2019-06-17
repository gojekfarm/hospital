package storage

import (
	"encoding/json"
)

// InsertOperation inserts operations
func InsertOperation(alertID int, jobName, script, status string) {
	sqlStatement := `INSERT INTO operations (surgeon_id, script, status, alert_id)
						VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, jobName, script, status, alertID)
	if err != nil {
		panic(err)
	}
}

// GetOperation returns the script.
func GetOperation(surgeonID string) (string, error) {
	rows, err := db.Query(
		`SELECT id, script FROM operations WHERE surgeon_id = $1 and status = $2`,
		surgeonID, "firing")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	ops := make([]*operation, 0)

	for rows.Next() {
		var (
			id     int
			script string
		)

		if err := rows.Scan(&id, &script); err != nil {
			return "", err
		}
		ops = append(ops, &operation{id, script})
	}

	jsonStr, err := json.Marshal(ops)
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}

type operation struct {
	ID     int    `json:"id"`
	Script string `json:"script"`
}
