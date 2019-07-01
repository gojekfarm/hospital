package dashboard

import (
	"fmt"
	"hospital/storage"
	"log"
	"net/http"
	"text/template"
)

// HandleDashboard for dashboard.
func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	mappings, err := storage.GetMappings()
	if err != nil {
		log.Println("GetMappings: " + err.Error())
	}
	t, _ := template.ParseFiles("dashboard/views/home.tpl", "dashboard/views/header.tpl", "dashboard/views/footer.tpl")

	resp := struct {
		Page string
		Maps []*storage.Mapping
	}{
		"mapping",
		mappings,
	}

	t.Execute(w, resp)
}

// HandleInsert for adding mapings.
func HandleInsert(w http.ResponseWriter, r *http.Request) {
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

// HandleRemove for removing a mapping.
func HandleRemove(w http.ResponseWriter, r *http.Request) {
	alertType := r.URL.Path[18:]

	err := storage.DeleteMapping(alertType)

	if err != nil {
		fmt.Fprintf(w, "Error deleting the mapping!")
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
