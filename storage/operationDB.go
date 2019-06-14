package storage

import (
	"encoding/json"
	"log"
)

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
		rows.Scan(&id, &script)
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
