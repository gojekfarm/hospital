package server

import (
	"net/http"
	"hospital/routes"
	"hospital/storage"
)



// StartServer will start the server
func StartServer(port string){
	storage.Connect()
	
	routes.Routes()

	http.ListenAndServe(":"+port, nil)
}