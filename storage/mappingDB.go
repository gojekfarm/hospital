package storage

import (
	"database/sql"
	"log"
)

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
		log.Println(err)
	}
	return script
}
