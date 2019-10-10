package message

import "net/http"

// HTTPRequest wraps http request and adds additional functions
type HTTPRequest struct {
	http.Request
	pathVariables map[string]string
}

// NewHTTPRequest is wrapper for original http request
func NewHTTPRequest(request *http.Request) *HTTPRequest {
	return &HTTPRequest{
		Request:       *request,
		pathVariables: make(map[string]string),
	}
}

// HeaderNames returns slice with all current header names
func (r HTTPRequest) HeaderNames() []string {
	headers := r.Header
	headerKeys := make([]string, 0, len(headers))
	for header := range headers {
		headerKeys = append(headerKeys, header)
	}
	return headerKeys
}

// GetPathVariables returnes resolved url variables
func (r *HTTPRequest) GetPathVariables() map[string]string {
	if r.pathVariables == nil {
		r.pathVariables = make(map[string]string)
	}
	return r.pathVariables
}

// SetPathVariables Replaces url variables with new map
func (r *HTTPRequest) SetPathVariables(pathVariables map[string]string) {
	if pathVariables == nil {
		pathVariables = make(map[string]string)
	}
	r.pathVariables = pathVariables
}
