package storage

import "database/sql"

// GetScript returns the script
func GetScript(alertType string) (string, error) {
	var script string
	err := db.QueryRow(`SELECT script FROM mapping WHERE alert_type = $1`,
		alertType).Scan(&script)

	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return script, err
}

// InsertScript adds a script to that alert_type.
func InsertScript(alertType, script string) error {
	_, err := GetScript(alertType)
	if err != nil {
		sqlStatement := `INSERT INTO mapping (alert_type, script)
						VALUES ($1, $2)`
		_, err = db.Exec(sqlStatement, alertType, script)
		return err
	}
	return err
}
