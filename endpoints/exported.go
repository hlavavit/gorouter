package endpoints

import (
	"github.com/hlavavit/gorouter/message"
	"github.com/hlavavit/gorouter/msghandlers"
	"github.com/hlavavit/gorouter/pathmatcher"
)

// Endpoint is interface with one method HandleMessage which receives pointers to both request and response.
// If message was proccessed by this enpoint, method must return true
type Endpoint interface {
	// HandleMessage when message was proccessed by this enpoint, method must return true
	HandleMessage(request *message.HTTPRequest, response *message.HTTPResponse) bool
}

// NewBasicEndpoint return implementation of endpoint interface using NewDefaultMatcher from package pathMatcher to determine if request should be processed by handler (second param)
// Basic endpoint also saves extracted path variables to request
func NewBasicEndpoint(pattern string, messageHandler msghandlers.MessageHandler, filters ...Filter) Endpoint {
	endpoint := &basic{
		matcher:        pathmatcher.NewDefaultMatcher(pattern),
		messageHandler: messageHandler,
		filters:        filters,
	}
	return endpoint
}

func (ep *basic) HandleMessage(request *message.HTTPRequest, response *message.HTTPResponse) bool {
	return ep.handleMessage(request, response)
}

//Filter is for preproccessing http request or post proccessing http response.
//should call proceed function before changing response,
//alternatively can not call proceed, effectively stopping request processing.
//In that case filter is responsible for writing response and setting response status
type Filter func(rq *message.HTTPRequest, rsp *message.HTTPResponse, proceed msghandlers.MessageHandler)
