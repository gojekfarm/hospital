package healthcheck

import (
	"fmt"
	"net/http"
)

// Handler gives response pong.
func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("received health-check request")
	fmt.Fprintf(w, "Pong")
}
