package doctor

import (
	"hospital/storage"
)

func scriptGenerator(alertType string) string {
	script := storage.GetScript(alertType)
	return script
}
