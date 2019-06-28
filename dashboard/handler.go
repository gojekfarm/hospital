package dashboard

import (
	"fmt"
	"hospital/storage"
	"log"
	"net/http"
	"text/template"
)

// Handler for dashboard.
func Handler(w http.ResponseWriter, r *http.Request) {
	mappings, err := storage.GetMappings()
	if err != nil {
		log.Println("GetMappings: " + err.Error())
	}
	t, _ := template.ParseFiles("dashboard/views/home.html", "dashboard/views/header.html", "dashboard/views/footer.html")

	resp := struct {
		Page string
		Maps []*storage.Mapping
	}{
		"active",
		mappings,
	}

	t.Execute(w, resp)
}

// InsertHandler for adding mapings.
func InsertHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.FormValue("alert") != "" && r.FormValue("script") != "" {
		err := storage.InsertScript(r.FormValue("alert"), r.FormValue("script"))
		if err != nil {
			fmt.Fprintf(w, "Script Adding Error!")
			return
		}
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

// RemoveHandler for removing a mapping.
func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	alertType := r.URL.Path[18:]

	err := storage.DeleteMapping(alertType)

	if err != nil {
		fmt.Fprintf(w, "Error deleting the mapping!")
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
