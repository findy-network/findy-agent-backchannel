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
	"encoding/json"
	"net/http"
	"strings"
)

// OutOfBandApiController binds http requests to an api service and writes the service results to the http response
type OutOfBandApiController struct {
	service      OutOfBandApiServicer
	errorHandler ErrorHandler
}

// OutOfBandApiOption for how the controller is set up.
type OutOfBandApiOption func(*OutOfBandApiController)

// WithOutOfBandApiErrorHandler inject ErrorHandler into controller
func WithOutOfBandApiErrorHandler(h ErrorHandler) OutOfBandApiOption {
	return func(c *OutOfBandApiController) {
		c.errorHandler = h
	}
}

// NewOutOfBandApiController creates a default api controller
func NewOutOfBandApiController(s OutOfBandApiServicer, opts ...OutOfBandApiOption) Router {
	controller := &OutOfBandApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the OutOfBandApiController
func (c *OutOfBandApiController) Routes() Routes {
	return Routes{
		{
			"OutOfBandCreateInvitation",
			strings.ToUpper("Post"),
			"/agent/command/out-of-band/send-invitation-message",
			c.OutOfBandCreateInvitation,
		},
		{
			"OutOfBandReceiveInvitation",
			strings.ToUpper("Post"),
			"/agent/command/out-of-band/{receive-invitation:receive-invitation\\/?}",
			c.OutOfBandReceiveInvitation,
		},
	}
}

// OutOfBandCreateInvitation - Create a new out of band invitation
func (c *OutOfBandApiController) OutOfBandCreateInvitation(w http.ResponseWriter, r *http.Request) {
	outOfBandCreateInvitationRequestParam := OutOfBandCreateInvitationRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&outOfBandCreateInvitationRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOutOfBandCreateInvitationRequestRequired(outOfBandCreateInvitationRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.OutOfBandCreateInvitation(r.Context(), outOfBandCreateInvitationRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// OutOfBandReceiveInvitation - Receive an out of band invitation
func (c *OutOfBandApiController) OutOfBandReceiveInvitation(w http.ResponseWriter, r *http.Request) {
	outOfBandReceiveInvitationRequestParam := OutOfBandReceiveInvitationRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&outOfBandReceiveInvitationRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOutOfBandReceiveInvitationRequestRequired(outOfBandReceiveInvitationRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.OutOfBandReceiveInvitation(r.Context(), outOfBandReceiveInvitationRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
