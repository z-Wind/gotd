package gotd

import (
	"context"
	"net/http"
	"net/url"
)

// ServerResponse is embedded in each Do response and
// provides the HTTP status code and header sent by the server.
type ServerResponse struct {
	// HTTPStatusCode is the server's response status code. When using a
	// resource method's Do call, this will always be in the 2xx range.
	HTTPStatusCode int
	// Header contains the response header fields from the server.
	Header http.Header
}

// DefaultCall call template
type DefaultCall struct {
	s         *Service
	urlParams url.Values
	ctx       context.Context
	header    http.Header
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultCall) Context(ctx context.Context) *DefaultCall {
	c.ctx = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultCall) Header() http.Header {
	if c.header == nil {
		c.header = make(http.Header)
	}
	return c.header
}
