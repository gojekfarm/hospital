package healthCheck

import (
	"fmt"
	"net/http"
)

// Handler gives response pong.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}
