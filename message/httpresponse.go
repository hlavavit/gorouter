package message

import (
	"io"
	"strings"
)

//HTTPResponse represents what will be written as response when proccessing ends
type HTTPResponse struct {
	Status HTTPStatus
	body   io.Reader
}

// NewHTTPResponse creates default ok emty http status
func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		Status: StatusOk,
	}
}

func (r HTTPResponse) GetBody() io.Reader {
	return r.body
}

func (r *HTTPResponse) SetStringBody(body string) {
	r.body = strings.NewReader(body)
}
