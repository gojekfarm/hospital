package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// func scriptHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		fmt.Fprintf(w, "Only post methods supported.")
// 	case "POST":
// 		body, _ := ioutil.ReadAll(r.Body)
// 		var respString AlertName
// 		json.Unmarshal(body, &respString)
// 		log.Println(respString.Alertname + " called!")

// 		db, err := sql.Open("mysql",
// 			"root:toor@tcp(127.0.0.1:3306)/Doctor")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer db.Close()
// 		var command string
// 		err = db.QueryRow("select script from Mapping where alert_type = ? ",
// 			respString.Alertname).Scan(&command)
// 		if err != nil {
// 			command = "none"
// 		}

// 		if command != "none" {
// 			log.Println(command)
// 			out, err := exec.Command("/bin/sh", "-c", command).Output()

// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			log.Println(string(out))
// 			fmt.Fprintf(w, "success")
// 		} else {
// 			fmt.Fprintf(w, "failed")
// 		}
// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		fmt.Fprintf(w, "I can't do that.")
// 	}
// }

func main() {
	// http.HandleFunc("/script", scriptHandler)

	// log.Println("Go!")
	// http.ListenAndServe(":8089", nil)

}

// type AlertName struct {
// 	Alertname string
// }