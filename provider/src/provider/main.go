package main

import (
	"net/http"
	"time"
	"fmt"
	"net"
	log "github.com/Sirupsen/logrus"
)

const port = "8080"

func main() {

	count := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.WithField("method", r.Method).WithField("uri", r.URL).WithField("count", count).Info("Serving request")
		text := fmt.Sprintf("Hello from the service provider. This is request number %d. The current time is %s.", count, time.Now().Local())
		w.Write([]byte(text))
		count++;
	})

	log.WithField("port", port).Info("Starting HTTP server")

	if err := http.ListenAndServe(net.JoinHostPort("", port), nil); err != nil {
		log.WithField("err", err).Fatal("Error setting up HTTP server")
	}

	log.Info("Exiting")
}


