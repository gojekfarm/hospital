package dashboard

import (
	"hospital/storage"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// HandleSummary handles summary.
func HandleSummary(w http.ResponseWriter, r *http.Request) {
	summaries, err := storage.GetSummary()
	if err != nil {
		log.Println("GetSummary: " + err.Error())
	}

	t, _ := template.ParseFiles("dashboard/views/summary.tpl", "dashboard/views/header.tpl", "dashboard/views/footer.tpl")

	resp := struct {
		Page      string
		Summaries map[string]*storage.Summary
	}{
		"summary",
		summaries,
	}

	err = t.Execute(w, resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// HandleOneSummary handle one summary
func HandleOneSummary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	summary, logs, err := storage.GetOneSummary(params["id"])
	if err != nil {
		log.Println("GetOneSummary: " + err.Error())
	}

	t, _ := template.ParseFiles("dashboard/views/singlesummary.tpl", "dashboard/views/header.tpl", "dashboard/views/footer.tpl")

	resp := struct {
		Page    string
		Summary storage.Summary
		Logs    []*storage.Logs
	}{
		"summary",
		summary,
		logs,
	}

	err = t.Execute(w, resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
