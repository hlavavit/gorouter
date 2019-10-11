package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hlavavit/gorouter/endpoints"
	"github.com/hlavavit/gorouter/message"

	"github.com/hlavavit/gorouter"
	log "github.com/sirupsen/logrus"
)

func initLogging() {
	log.SetOutput(os.Stdout)
	//log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
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

	endpoint := endpoints.NewBasicEndpoint("/test/{var}", func(rq *message.HTTPRequest, rsp *message.HTTPResponse) {
		log.Info("Received path variables", rq.GetPathVariables())
		rsp.Status = message.HTTPStatus(http.StatusOK)
		rsp.SetStringBody(fmt.Sprintf("path variables: %v", rq.GetPathVariables()))
	})
	teapod := endpoints.NewBasicEndpoint("/teapod", func(rq *message.HTTPRequest, rsp *message.HTTPResponse) {
		rsp.Status = message.HTTPStatus(http.StatusTeapot)
		rsp.SetStringBody(getTeapod())
	})

	router.AddEndpoints(endpoint, teapod)
	server := http.Server{Addr: ":8080", Handler: router}
	err <- server.ListenAndServe()
}

func getTeapod() (teapod string) {
	teapod += "src http://ascii.co.uk/art/teapot\n\n"
	teapod += "                       (\n"
	teapod += "            _           ) )\n"
	teapod += "         _,(_)._        ((\n"
	teapod += "    ___,(_______).        )\n"
	teapod += "  ,'__.   /       \\    /\\_\n"
	teapod += " /,' /  |\"\"|       \\  /  /\n"
	teapod += "| | |   |__|       |,'  /\n"
	teapod += " \\`.|                  /\n"
	teapod += "  `. :           :    /\n"
	teapod += "    `.            :.,'\n"
	teapod += "      `-.________,-'	"
	return teapod
}
