package gorouter

import (
	"github.com/hlavavit/gorouter/endpoints"
	"github.com/hlavavit/gorouter/msghandlers"
)

type router struct {
	endpoints             []endpoints.Endpoint
	defaultMessageHandler msghandlers.MessageHandler
}

func (r *router) addEndpoints(added ...endpoints.Endpoint) {
	r.endpoints = append(r.endpoints, added...)
}
