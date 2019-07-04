package storage

// InsertOperation inserts operations
func InsertOperation(alertID int, applicationID, script, status string) int {
	lastInsertID := 0
	sqlStatement := `INSERT INTO operations (application_id, script, status, alert_id)
						VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(sqlStatement, applicationID, script, status, alertID).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}

	return lastInsertID
}

// GetOperation returns the script.
func GetOperation(applicationID string) ([]*Operation, error) {
	ops := make([]*Operation, 0)

	rows, err := db.Query(
		`SELECT id, script FROM operations WHERE application_id = $1 and status = $2`,
		applicationID, "CRITICAL")
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

// GetApplicationID returns the applicationID of corresponding ID.
func GetApplicationID(id int) (string, error) {
	var applicationID string

	err := db.QueryRow(`SELECT application_id FROM operations WHERE id = $1`,
		id).Scan(&applicationID)

	return applicationID, err
}

// GetLogs give list of logs present in the table.
func GetLogs() ([]*Logs, error) {
	logs := make([]*Logs, 0)

	rows, err := db.Query(
		`SELECT  application_id, script, status, logs FROM operations`)
	if err != nil {
		return logs, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			applicationID string
			script        string
			status        string
			log           string
		)

		if err := rows.Scan(&applicationID, &script, &status, &log); err != nil {
			return logs, err
		}
		logs = append(logs, &Logs{applicationID, script, status, log})
	}

	return logs, nil
}

// Logs struct for getting all entries in table.
type Logs struct {
	ApplicationID string
	Script        string
	Status        string
	Logs          string
}
