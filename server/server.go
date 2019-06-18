package server

import (
	"fmt"
	"hospital/routes"
	"hospital/storage"
	"net/http"
	"os"
)

// StartServer will start the server
func StartServer() {
	storage.Initialize()

	routes.Routes()

	err := http.ListenAndServe(os.Getenv("HOST_ADDRESS")+":"+os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Println(err)
	}
}
