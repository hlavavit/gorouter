package msghandlers

import "github.com/hlavavit/gorouter/message"

// MessageHandler type for message handler functions 
type MessageHandler func(rq *message.HTTPRequest, rsp *message.HTTPResponse)

// HandleNotFound return status 404 with simple not found page
func HandleNotFound(rq *message.HTTPRequest, rsp *message.HTTPResponse) {
	handleNotFound(rq, rsp)
}
