package dashboard

import (
	"hospital/storage"
	"log"
	"net/http"
	"text/template"
)

// HandleSummary handles summary.
func HandleSummary(w http.ResponseWriter, r *http.Request) {
	mappings, err := storage.GetMappings()
	if err != nil {
		log.Println("GetMappings: " + err.Error())
	}
	t, _ := template.ParseFiles("dashboard/views/summary.tpl", "dashboard/views/header.tpl", "dashboard/views/footer.tpl")

	resp := struct {
		Page string
		Maps []*storage.Mapping
	}{
		"summary",
		mappings,
	}

	t.Execute(w, resp)
}
