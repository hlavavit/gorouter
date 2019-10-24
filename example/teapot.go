package main

import (
	"net/http"

	"github.com/hlavavit/gorouter/message"
)

// HandleTeapot is menthod for handling http request, writes nice image of teapot as response
func HandleTeapot(rq *message.HTTPRequest, rsp *message.HTTPResponse) {
	rsp.Status = message.HTTPStatus(http.StatusTeapot)
	rsp.SetStringBody(getTeapod())

}

func getTeapod() (teapod string) {
	teapod += "src http://ascii.co.uk/art/teapot\n\n"
	teapod += "                       (\n"
	teapod += "            _           ) )\n"
	teapod += "         _,(_)._        ((\n"
	teapod += "    ___,(_______).        )\n"
	teapod += "  ,'__.   /       \\    /\\_\n"
	teapod += " /,' /  |\"\"|       \\  /  /\n"
	teapod += "| | |   |__|       |,'  /\n"
	teapod += " \\`.|                  /\n"
	teapod += "  `. :           :    /\n"
	teapod += "    `.            :.,'\n"
	teapod += "      `-.________,-'	"
	return teapod
}
