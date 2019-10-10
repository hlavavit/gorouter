package gorouter

import "github.com/hlavavit/gorouter/endpoints"

type Router struct {
	endpoints []endpoints.Endpoint
}

func (r *Router) AddEndpoints(added ...endpoints.Endpoint) {
	r.endpoints = append(r.endpoints, added...)
}
