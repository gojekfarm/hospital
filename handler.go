package main

import (
	"fmt"
	"net/http"
)

// PingHandler gives response pong.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}
