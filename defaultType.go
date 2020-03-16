package gotd

import (
	"context"
	"net/http"
	"net/url"
)

// DefaultCall DefaultCall function
type DefaultCall struct {
	s         *Service
	urlParams url.Values
	ctx       context.Context
	header    http.Header
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsGetAccountCall) Context(ctx context.Context) *AccountsGetAccountCall {
	c.ctx = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsGetAccountCall) Header() http.Header {
	if c.header == nil {
		c.header = make(http.Header)
	}
	return c.header
}
