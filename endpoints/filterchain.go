package endpoints

import (
	"github.com/hlavavit/gorouter/message"
	"github.com/hlavavit/gorouter/msghandlers"
)

func (ep *basic) processFilterchain(request *message.HTTPRequest, response *message.HTTPResponse) {
	if ep.filters != nil {
		count := len(ep.filters)
		idx := 0

		var proceed msghandlers.MessageHandler
		proceed = func(rq *message.HTTPRequest, rsp *message.HTTPResponse) {
			if idx < count {
				filter := ep.filters[idx]
				idx++
				filter(rq, rsp, proceed)
			} else {
				ep.messageHandler(rq, rsp)
			}
		}
		proceed(request, response)

	} else {
		ep.messageHandler(request, response)
	}
}
