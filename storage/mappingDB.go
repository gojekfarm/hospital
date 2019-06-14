package storage

import "database/sql"

// GetScript returns the script
func GetScript(alertType string) string {
	var script string
	err := db.QueryRow(`SELECT script FROM mapping WHERE alert_type = $1`,
		alertType).Scan(&script)

	if err == sql.ErrNoRows {
		//log.Fatal("No Results Found")
		return "no script"
	}

	if err != nil {
		//log.Fatal(err)
	}
	return script
}
