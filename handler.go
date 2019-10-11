package gorouter

import (
	"io"

	"github.com/hlavavit/gorouter/message"
	log "github.com/sirupsen/logrus"

	"net/http"
)

func (r router) serveHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	//log.Info("Start")

	httpRequest := message.NewHTTPRequest(request)
	httpResponse := message.NewHTTPResponse()
	processed := false
	for _, endpoint := range r.endpoints {
		if processed = endpoint.HandleMessage(httpRequest, httpResponse); processed {
			break
		}
	}

	if !processed {
		r.defaultMessageHandler(httpRequest, httpResponse)
	}

	//responseWriter.Header().Add("test-header", "test-value")
	responseWriter.WriteHeader(int(httpResponse.Status))
	if httpResponse.GetBody() != nil {
		_, err := io.Copy(responseWriter, httpResponse.GetBody())
		if err != nil {
			log.Error("Failed to copy content to response", err)
		}
	}

	//log.Info("Exit")
}
