package dashboard

import (
	"fmt"
	"hospital/storage"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
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

	err = t.Execute(w, resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// HandleInsert for adding mapings.
func HandleInsert(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

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
	params := mux.Vars(r)
	alertType := params["alertType"]

	err := storage.DeleteMapping(alertType)

	if err != nil {
		fmt.Fprintf(w, "Error deleting the mapping!")
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
