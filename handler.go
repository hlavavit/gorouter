package gorouter

import (
	"github.com/hlavavit/gorouter/message"

	"net/http"
)

func (r Router) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	//log.Info("Start")

	httpRequest := message.NewHTTPRequest(request)
	httpResponse := message.NewHTTPResponse()

	for _, endpoint := range r.endpoints {
		if endpoint.HandleMessage(httpRequest, httpResponse) {
			break
		}
	}

	//responseWriter.Header().Add("test-header", "test-value")
	responseWriter.WriteHeader(int(httpResponse.Status))
	//responseWriter.Write([]byte("testik"))

	//log.Info("Exit")
}
