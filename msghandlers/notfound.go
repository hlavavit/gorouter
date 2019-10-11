package msghandlers

import (
	"github.com/hlavavit/gorouter/message"
	log "github.com/sirupsen/logrus"
)

var notFound = `
<!DOCTYPE html>
<html>
  <body>
    <h1>404 - Not Found</h1>
  </body>
</html>
`

// HandleNotFound return status 404 with simple not found page
func handleNotFound(rq *message.HTTPRequest, rsp *message.HTTPResponse) {
	log.Tracef("returning 404 for %v %v", rq.Method, rq.RequestURI)
	rsp.Status = message.StatusNotFound
	rsp.SetStringBody(notFound)
}
