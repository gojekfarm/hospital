package reception

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../dbprovider"
)

// ReceptionHandler recieves alerts
func ReceptionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		var receivedObject received
		json.Unmarshal(body, &receivedObject)
		resp := alertReceiver(receivedObject)
		fmt.Fprintf(w, resp)
	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}

func alertReceiver(receivedObj received) string {
	dbprovider.InsertAlertUnique(receivedObj.Alerts[0].Labels.Alertname, receivedObj.Alerts[0].StartsAt, receivedObj.Alerts[0].Labels.Instance, receivedObj.Alerts[0].Status)
	return "success"
}

// func acceptHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {
// 	case "GET":
// 		fmt.Fprintf(w, "Only post methods supported.")
// 	case "POST":
// 		var receivedObject received
// 		body, _ := ioutil.ReadAll(r.Body)
// 		json.Unmarshal(body, &receivedObject)
// 		db, err := sql.Open("mysql",
// 			"root:toor@tcp(127.0.0.1:3306)/Doctor")
// 		if err != nil {r
// 			log.Fatal(err)
// 		}
// 		defer db.Close()

// 		if receivedObject.Alerts[0].Status == "resolved" {
// 			log.Println("resolved condition called!")
// 			var id int
// 			err = db.QueryRow("select id from Incidents where alertname = ? and startsAT = ? and address = ?",
// 				receivedObject.Alerts[0].Labels.Alertname, receivedObject.Alerts[0].StartsAt, receivedObject.Alerts[0].Labels.Instance).Scan(&id)
// 			if err != nil {

// 			}

// 			insert, err := db.Query("UPDATE Incidents set status = ? WHERE id = ?",
// 				receivedObject.Alerts[0].Status, id)

// 			// if there is an error inserting, handle it
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			// be careful deferring Queries if you are using transactions
// 			defer insert.Close()
// 		} else {
// 			insert, err := db.Query("INSERT INTO Incidents(alertname, startsAT, address, status) VALUES ( ? , ? , ? ,?)",
// 				receivedObject.Alerts[0].Labels.Alertname, receivedObject.Alerts[0].StartsAt, receivedObject.Alerts[0].Labels.Instance, receivedObject.Alerts[0].Status)

// 			// if there is an error inserting, handle it
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			// be careful deferring Queries if you are using transactions
// 			defer insert.Close()

// 			var jsonStr = []byte(`{"alertname":"` + receivedObject.Alerts[0].Labels.Alertname + `"}`)
// 			req, err := http.NewRequest("POST", "http://localhost:8089/script", bytes.NewBuffer(jsonStr))
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			client := &http.Client{}
// 			resp, err := client.Do(req)
// 			if err != nil {
// 				panic(err)
// 			}
// 			body, _ := ioutil.ReadAll(resp.Body)
// 			log.Println(body)
// 			defer resp.Body.Close()

// 		}

// 		fmt.Fprintf(w, "success")

// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		fmt.Fprintf(w, "I can't do that.")
// 	}
// }

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
