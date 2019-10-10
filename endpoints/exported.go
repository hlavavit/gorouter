package endpoints

import (
	"github.com/hlavavit/gorouter/message"
	"github.com/hlavavit/gorouter/pathmatcher"
)

// Endpoint is interface with one method HandleMessage which receives pointers to both request and response.
// If message was proccessed by this enpoint, method should return true
type Endpoint interface {
	HandleMessage(request *message.HTTPRequest, response *message.HTTPResponse) bool
}

// NewBasicEndpoint return implementation of endpoint interface using NewDefaultMatcher from package pathMatcher to determine if request should be processed by handler (second param)
// Basic endpoint also saves extracted path variables to request
func NewBasicEndpoint(pattern string, messageHandler func(*message.HTTPRequest, *message.HTTPResponse)) Endpoint {
	endpoint := &basic{
		matcher:        pathmatcher.NewDefaultMatcher(pattern),
		messageHandler: messageHandler,
	}
	return endpoint
}

func (ep *basic) HandleMessage(request *message.HTTPRequest, response *message.HTTPResponse) bool {
	return ep.handleMessage(request, response)
}
