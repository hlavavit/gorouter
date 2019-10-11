package gorouter

import (
	"net/http"

	"github.com/hlavavit/gorouter/msghandlers"

	"github.com/hlavavit/gorouter/endpoints"
)

// Router wraps request and prepares empty response, delegates to added endpoints.
// Uses first enpoint that returns true (processed) and skips rest.
// When no endpoint matches uses default endpoint - 404
type Router interface {
	http.Handler
	// Add endpoints considered for proccessing. Order matters
	AddEndpoints(added ...endpoints.Endpoint)
	// Endpoids should have 404 default, this method is just for changing it
	SetDefaultMessageHandler(mh msghandlers.MessageHandler)
}

// New creates new instance of router with default 404 message nadler
func New() Router {
	router := &router{
		defaultMessageHandler: msghandlers.HandleNotFound,
	}
	return router
}

func (r *router) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	r.serveHTTP(rw, rq)
}

func (r *router) AddEndpoints(added ...endpoints.Endpoint) {
	r.addEndpoints(added...)
}

func (r *router) SetDefaultMessageHandler(mh msghandlers.MessageHandler) {
	r.defaultMessageHandler = mh
}
