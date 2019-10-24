package main

import (
	"net/http"
	"os"

	"github.com/hlavavit/gorouter/endpoints"

	"github.com/hlavavit/gorouter"
	log "github.com/sirupsen/logrus"
)

func initLogging() {
	log.SetOutput(os.Stdout)
	//log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
	log.SetLevel(log.TraceLevel)
}

func main() {
	initLogging()
	log.Info("Starting with arguments", os.Args)

	err := make(chan error)
	go run(err)

	if <-err != nil {
		log.Error("Error while starting a server", err)

		os.Exit(1)
	}
	log.Info("Shutdown...")

}

func run(err chan error) {
	router := gorouter.New()

	variable := endpoints.NewBasicEndpoint("/variable/{var}", HandleVariables)
	teapod := endpoints.NewBasicEndpoint("/teapot", HandleTeapot)

	router.AddEndpoints(variable, teapod)
	server := http.Server{Addr: ":8080", Handler: router}
	err <- server.ListenAndServe()
}
