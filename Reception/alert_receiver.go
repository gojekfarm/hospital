package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func acceptHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Only post methods supported.")
	case "POST":
		filename := "alert_api.txt"
		body, _ := ioutil.ReadAll(r.Body)
		ioutil.WriteFile(filename, body, 0600)
		var receivedObject received
		json.Unmarshal(body, &receivedObject)
		db, err := sql.Open("mysql",
			"root:toor@tcp(127.0.0.1:3306)/Doctor")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if receivedObject.Alerts[0].Status == "resolved" {
			log.Println("resolved condition called!")
			var id int
			err = db.QueryRow("select id from Incidents where alertname = ? and startsAT = ? and address = ?",
				receivedObject.Alerts[0].Labels.Alertname, receivedObject.Alerts[0].StartsAt, receivedObject.Alerts[0].Labels.Instance).Scan(&id)
			if err != nil {

			}

			insert, err := db.Query("UPDATE Incidents set status = ? WHERE id = ?",
				receivedObject.Alerts[0].Status, id)

			// if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}
			// be careful deferring Queries if you are using transactions
			defer insert.Close()
		} else {
			insert, err := db.Query("INSERT INTO Incidents(alertname, startsAT, address, status) VALUES ( ? , ? , ? ,?)",
				receivedObject.Alerts[0].Labels.Alertname, receivedObject.Alerts[0].StartsAt, receivedObject.Alerts[0].Labels.Instance, receivedObject.Alerts[0].Status)

			// if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}
			// be careful deferring Queries if you are using transactions
			defer insert.Close()

			var jsonStr = []byte(`{"alertname":"` + receivedObject.Alerts[0].Labels.Alertname + `"}`)
			req, err := http.NewRequest("POST", "http://localhost:8089/script", bytes.NewBuffer(jsonStr))
			if err != nil {
				log.Fatal(err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println(body)
			defer resp.Body.Close()

		}

		fmt.Fprintf(w, "success")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	filename := "alert_api.txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(w, "error")
	}
	var receivedObject received
	json.Unmarshal(body, &receivedObject)
	fmt.Fprintf(w, receivedObject.Alerts[0].Labels.Alertname)
}

func main() {
	http.HandleFunc("/accept", acceptHandler)
	http.HandleFunc("/view", viewHandler)

	log.Println("Go!")
	http.ListenAndServe(":8088", nil)
}

type received struct {
	Receiver          string
	Status            string
	Alerts            []alert
	GroupLabels       string
	CommonLabels      string
	CommonAnnotations string
	ExternalURL       string
	Version           string
	GroupKey          string
}

type alert struct {
	Status       string
	Labels       label
	Annotations  string
	StartsAt     string
	EndsAt       string
	GeneratorURL string
}

type label struct {
	Alertname string
	Backend   string
	Instance  string
	Job       string
	Severity  string
}
