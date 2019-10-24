package main

import (
	"fmt"
	"net/http"

	"github.com/hlavavit/gorouter/message"
	log "github.com/sirupsen/logrus"
)

//HandleVariables will write all path variables, params and body content to response.
func HandleVariables(rq *message.HTTPRequest, rsp *message.HTTPResponse) {
	log.Info("Received path variables", rq.GetPathVariables())
	rsp.Status = message.HTTPStatus(http.StatusOK)
	rsp.SetStringBody(fmt.Sprintf("path variables: %v", rq.GetPathVariables()))
	//TODO right now only handles path variables
}
