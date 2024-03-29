/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"
)

// StatusApiController binds http requests to an api service and writes the service results to the http response
type StatusApiController struct {
	service      StatusApiServicer
	errorHandler ErrorHandler
}

// StatusApiOption for how the controller is set up.
type StatusApiOption func(*StatusApiController)

// WithStatusApiErrorHandler inject ErrorHandler into controller
func WithStatusApiErrorHandler(h ErrorHandler) StatusApiOption {
	return func(c *StatusApiController) {
		c.errorHandler = h
	}
}

// NewStatusApiController creates a default api controller
func NewStatusApiController(s StatusApiServicer, opts ...StatusApiOption) Router {
	controller := &StatusApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the StatusApiController
func (c *StatusApiController) Routes() Routes {
	return Routes{
		{
			"StatusGet",
			strings.ToUpper("Get"),
			"/agent/command/{status:status\\/?}",
			c.StatusGet,
		},
		{
			"VersionGet",
			strings.ToUpper("Get"),
			"/agent/command/{version:version\\/?}",
			c.VersionGet,
		},
	}
}

// StatusGet - Get agent/backchannel status
func (c *StatusApiController) StatusGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.StatusGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// VersionGet - Get agent/backchannel version
func (c *StatusApiController) VersionGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.VersionGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
