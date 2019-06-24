package surgeon

import (
	"hospital/surgeon/backoff"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	applicationID = os.Getenv("APPLICATION_ID")
	url           = os.Getenv("HOST_PROTOCOL") + "://" + os.Getenv("HOST_ADDRESS") +
		":" + os.Getenv("PORT")
)

// LongPolling will do long polling.
func LongPolling() {
	maxWait, err := strconv.Atoi(os.Getenv("POLLING_WAIT_SECONDS"))
	if err != nil {
		panic(err)
	}

	maxBackoff, err := strconv.Atoi(os.Getenv("MAX_EXPONENTIAL_WAIT"))
	if err != nil {
		panic(err)
	}

	b := &backoff.Backoff{
		//These are the defaults
		Min:    2 * time.Duration(maxWait) * time.Second,
		Max:    time.Duration(maxBackoff) * time.Second,
		Factor: 2,
		Jitter: true,
	}

	for {
		err := makeRequest()

		rand.Seed(time.Now().UTC().UnixNano())

		if err == nil {
			waitTime := 1 + rand.Intn(maxWait-1)
			time.Sleep(time.Duration(waitTime) * time.Second)
			b.Reset()
		} else {
			d := b.Duration()
			log.Println(err.Error() + " reconnecting in " + d.String())
			time.Sleep(d)
		}

	}

}
