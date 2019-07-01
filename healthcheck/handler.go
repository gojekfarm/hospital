package healthcheck

import (
	"fmt"
	"net/http"
)

// Handler gives response pong.
func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}
