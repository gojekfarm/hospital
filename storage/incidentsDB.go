package storage

import "database/sql"

// InsertAlert insert alert into DB
func InsertAlert(alertname, startsAT, address, status string) {
	sqlStatement := `INSERT INTO incidents (alertname, starts_at, address, status)
						VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, alertname, startsAT, address, status)
	if err != nil {
		panic(err)
	}

}

// InsertAlertUnique insert alert into DB unique
func InsertAlertUnique(alertname, startsAT, address, status string) {
	id := GetAlertID(alertname, startsAT, address)
	if id != -1 {
		sqlStatement := `UPDATE incidents SET status = $2 WHERE id = $1;`
		_, err := db.Exec(sqlStatement, id, status)
		if err != nil {
			panic(err)
		}

	} else {
		InsertAlert(alertname, startsAT, address, status)
	}

}

// GetAlertID returns the alert id
func GetAlertID(alertname, startsAT, address string) int {
	var id int

	err := db.QueryRow(`SELECT id FROM incidents WHERE alertname = $1 and starts_at = $2 and address = $3`,
		alertname, startsAT, address).Scan(&id)

	if err == sql.ErrNoRows {
		//log.Fatal("No Results Found")
		id = -1
	}

	if err != nil {
		//log.Fatal(err)
	}
	return id
}
