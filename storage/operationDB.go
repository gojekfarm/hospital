package storage

// InsertOperation inserts operations
func InsertOperation(alertID int, applicationID, script, status string) {
	sqlStatement := `INSERT INTO operations (surgeon_id, script, status, alert_id)
						VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, applicationID, script, status, alertID)
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

// RecordStatus changes status of an operation.
func RecordStatus(id int, status, logs string) error {
	sqlStatement := `UPDATE operations SET status = $2, logs = $3 WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id, status, logs)
	if err != nil {
		return err
	}
	return nil
}

// AlertNameFromOpID returns alertname from operation ID.
func AlertNameFromOpID(id int) (string, error) {
	var alertID int

	err := db.QueryRow(`SELECT alert_id FROM operations WHERE id = $1`,
		id).Scan(&alertID)

	if err != nil {
		return "", err
	}

	alertName, err := GetAlertName(alertID)

	return alertName, err
}

// GetSurgeonID returns the surgeonID of corresponding ID.
func GetSurgeonID(id int) (string, error) {
	var surgeonID string

	err := db.QueryRow(`SELECT surgeon_id FROM operations WHERE id = $1`,
		id).Scan(&surgeonID)

	if err != nil {
		return "", err
	}

	return surgeonID, err
}
