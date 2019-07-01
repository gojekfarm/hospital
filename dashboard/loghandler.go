package dashboard

import (
	"hospital/storage"
	"log"
	"net/http"
	"text/template"
)

// HandleLogs for dashboard.
func HandleLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := storage.GetLogs()
	if err != nil {
		log.Println("GetLogs: " + err.Error())
	}
	t, _ := template.ParseFiles("dashboard/views/logs.tpl", "dashboard/views/header.tpl", "dashboard/views/footer.tpl")

	resp := struct {
		Page string
		Logs []*storage.Logs
	}{
		"logs",
		logs,
	}

	t.Execute(w, resp)
}
