package storage

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
func GetOperation(surgeonID string) ([]*Operation, error) {
	ops := make([]*Operation, 0)

	rows, err := db.Query(
		`SELECT id, script FROM operations WHERE surgeon_id = $1 and status = $2`,
		surgeonID, "firing")
	if err != nil {
		return ops, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id     int
			script string
		)

		if err := rows.Scan(&id, &script); err != nil {
			return ops, err
		}
		ops = append(ops, &Operation{id, script})
	}

	return ops, nil
}

// Operation struct for response to surgeon.
type Operation struct {
	ID     int    `json:"id"`
	Script string `json:"script"`
}

// ReportStatus changes status of an operation.
func ReportStatus(id int, status, logs string) error {
	sqlStatement := `UPDATE operations SET status = $2, logs = $3 WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id, status, logs)
	if err != nil {
		return err
	}
	return nil
}
