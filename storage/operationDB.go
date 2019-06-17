package storage

import (
	"encoding/json"
	"log"
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

// GetOperation returns the script
func GetOperation(surgeonID string) string {
	rows, err := db.Query(`SELECT id, script FROM operations WHERE surgeon_id = $1 and status = $2`,
		surgeonID, "firing")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var id int
	var script string
	var oparray []operation
	for rows.Next() {
		err := rows.Scan(&id, &script)
		if err != nil {
			log.Println(err)
		}
		oparray = append(oparray, operation{id, script})
	}
	jsonStr, err := json.Marshal(oparray)
	if err != nil {
		log.Println(err)
	}

	return string(jsonStr)
}

type operation struct {
	ID     int    `json:"id"`
	Script string `json:"script"`
}
