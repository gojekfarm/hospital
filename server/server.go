package server

import (
	"hospital/routes"
	"hospital/storage"
	"net/http"
	"os"
)

// StartServer will start the server
func StartServer() {
	storage.Initialize()

	routes.Routes()

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
