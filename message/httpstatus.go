package message

//HTTPStatus provides convinience functions for resolving status codes
type HTTPStatus int

// common status codes
const (
	StatusOk       HTTPStatus = 200
	StatusNotFound HTTPStatus = 404
)
