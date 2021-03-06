package storage

import (
	"database/sql"
	"log"
)

// InsertAlert insert alert into DB
func InsertAlert(alertname, startsAT, applicationID, status string) int {
	lastInsertID := 0
	sqlStatement := `INSERT INTO incidents (alertname, starts_at, application_id, status)
						VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(sqlStatement, alertname, startsAT, applicationID, status).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	return lastInsertID
}

// InsertAlertUnique insert alert into DB unique
func InsertAlertUnique(alertname, startsAT, applicationID, status string) int {
	id := GetAlertID(alertname, startsAT, applicationID)
	if id != -1 {
		sqlStatement := `UPDATE incidents SET status = $2 WHERE id = $1;`
		_, err := db.Exec(sqlStatement, id, status)
		if err != nil {
			panic(err)
		}

	} else {
		id = InsertAlert(alertname, startsAT, applicationID, status)

	}
	return id
}

// GetAlertID returns the alert id
func GetAlertID(alertname, startsAT, applicationID string) int {
	var id int

	err := db.QueryRow(`SELECT id FROM incidents WHERE alertname = $1 and starts_at = $2 and application_id = $3`,
		alertname, startsAT, applicationID).Scan(&id)

	if err == sql.ErrNoRows {
		//log.Fatal("No Results Found")
		id = -1
	}

	if err != nil {
		log.Println(err)
	}
	return id
}

// GetAlertName returns alertname of given ID
func GetAlertName(id int) (string, error) {
	var alertname string

	err := db.QueryRow(`SELECT alertname FROM incidents WHERE id = $1`,
		id).Scan(&alertname)

	return alertname, err
}
