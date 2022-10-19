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

// PresentProofV2ApiController binds http requests to an api service and writes the service results to the http response
type PresentProofV2ApiController struct {
	service      PresentProofV2ApiServicer
	errorHandler ErrorHandler
}

// PresentProofV2ApiOption for how the controller is set up.
type PresentProofV2ApiOption func(*PresentProofV2ApiController)

// WithPresentProofV2ApiErrorHandler inject ErrorHandler into controller
func WithPresentProofV2ApiErrorHandler(h ErrorHandler) PresentProofV2ApiOption {
	return func(c *PresentProofV2ApiController) {
		c.errorHandler = h
	}
}

// NewPresentProofV2ApiController creates a default api controller
func NewPresentProofV2ApiController(s PresentProofV2ApiServicer, opts ...PresentProofV2ApiOption) Router {
	controller := &PresentProofV2ApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PresentProofV2ApiController
func (c *PresentProofV2ApiController) Routes() Routes {
	return Routes{
		{
			"PresentProofV2SendPresentation",
			strings.ToUpper("Post"),
			"/agent/command/proof-v2/send-presentation",
			c.PresentProofV2SendPresentation,
		},
		{
			"PresentProofV2SendRequest",
			strings.ToUpper("Post"),
			"/agent/command/proof-v2/send-request",
			c.PresentProofV2SendRequest,
		},
		{
			"PresentProofV2VerifyPresentation",
			strings.ToUpper("Post"),
			"/agent/command/proof-v2/verify-presentation",
			c.PresentProofV2VerifyPresentation,
		},
	}
}

// PresentProofV2SendPresentation - Send presentation
func (c *PresentProofV2ApiController) PresentProofV2SendPresentation(w http.ResponseWriter, r *http.Request) {
	presentProofV2SendPresentationRequestParam := PresentProofV2SendPresentationRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&presentProofV2SendPresentationRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPresentProofV2SendPresentationRequestRequired(presentProofV2SendPresentationRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PresentProofV2SendPresentation(r.Context(), presentProofV2SendPresentationRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PresentProofV2SendRequest - Send presentation request
func (c *PresentProofV2ApiController) PresentProofV2SendRequest(w http.ResponseWriter, r *http.Request) {
	presentProofV2SendRequestRequestParam := PresentProofV2SendRequestRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&presentProofV2SendRequestRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPresentProofV2SendRequestRequestRequired(presentProofV2SendRequestRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PresentProofV2SendRequest(r.Context(), presentProofV2SendRequestRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PresentProofV2VerifyPresentation - Verify presentation
func (c *PresentProofV2ApiController) PresentProofV2VerifyPresentation(w http.ResponseWriter, r *http.Request) {
	presentProofVerifyPresentationRequestParam := PresentProofVerifyPresentationRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&presentProofVerifyPresentationRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPresentProofVerifyPresentationRequestRequired(presentProofVerifyPresentationRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PresentProofV2VerifyPresentation(r.Context(), presentProofVerifyPresentationRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
