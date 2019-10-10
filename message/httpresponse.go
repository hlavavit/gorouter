package message

//HTTPResponse represents what will be written as response when proccessing ends
type HTTPResponse struct {
	Status HTTPStatus
}

// NewHTTPResponse creates default ok emty http status
func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		Status: StatusOk,
	}
}
