package endpoints

import (
	"github.com/hlavavit/gorouter/message"
	"github.com/hlavavit/gorouter/msghandlers"
	"github.com/hlavavit/gorouter/pathmatcher"
)

type basic struct {
	matcher        pathmatcher.Matcher
	filters        []Filter
	messageHandler msghandlers.MessageHandler
}

func (ep basic) match(request *message.HTTPRequest) (match bool, pathVariables map[string]string) {
	pathVariables = make(map[string]string)
	if ep.matcher == nil {
		match = true
		return
	}
	return ep.matcher.Match(request.URL.Path)
}

func (ep basic) handleMessage(request *message.HTTPRequest, response *message.HTTPResponse) bool {
	match, pathVariables := ep.match(request)
	if !match {
		return false
	}
	request.SetPathVariables(pathVariables)
	ep.processFilterchain(request, response)

	return true
}
