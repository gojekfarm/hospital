package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Pong",name)
	time.Sleep(1 * time.Second)
}

func main() {
	http.HandleFunc("/ping/", PingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}
