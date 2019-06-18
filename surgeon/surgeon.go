package surgeon

import (
	"hospital/routes"
	"os"
	"strconv"
	"time"
)

var surgeonID = os.Getenv("SURGEON_ID")
var url = os.Getenv("HOST_PROTOCOL") + "://" + os.Getenv("HOST_ADDRESS") +
	":" + os.Getenv("PORT") + routes.OperationAPIPath

// LongPolling will do ling polling.
func LongPolling() {

	for {
		makeRequest()

		waitTime, err := strconv.Atoi(os.Getenv("POLLING_WAIT_SECONDS"))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Duration(waitTime) * time.Second)
	}

}
