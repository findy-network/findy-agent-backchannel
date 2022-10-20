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

// IssueCredentialV3ApiController binds http requests to an api service and writes the service results to the http response
type IssueCredentialV3ApiController struct {
	service      IssueCredentialV3ApiServicer
	errorHandler ErrorHandler
}

// IssueCredentialV3ApiOption for how the controller is set up.
type IssueCredentialV3ApiOption func(*IssueCredentialV3ApiController)

// WithIssueCredentialV3ApiErrorHandler inject ErrorHandler into controller
func WithIssueCredentialV3ApiErrorHandler(h ErrorHandler) IssueCredentialV3ApiOption {
	return func(c *IssueCredentialV3ApiController) {
		c.errorHandler = h
	}
}

// NewIssueCredentialV3ApiController creates a default api controller
func NewIssueCredentialV3ApiController(s IssueCredentialV3ApiServicer, opts ...IssueCredentialV3ApiOption) Router {
	controller := &IssueCredentialV3ApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IssueCredentialV3ApiController
func (c *IssueCredentialV3ApiController) Routes() Routes {
	return Routes{
		{
			"IssueCredentialV3Issue",
			strings.ToUpper("Post"),
			"/agent/command/issue-credential-v3/issue",
			c.IssueCredentialV3Issue,
		},
		{
			"IssueCredentialV3RetrieveCredentialApplication",
			strings.ToUpper("Get"),
			"/agent/command/issue-credential-v3/retrieve-credential-application",
			c.IssueCredentialV3RetrieveCredentialApplication,
		},
		{
			"IssueCredentialV3RetrieveCredentialProposal",
			strings.ToUpper("Get"),
			"/agent/command/issue-credential-v3/retrieve-credential-proposal",
			c.IssueCredentialV3RetrieveCredentialProposal,
		},
		{
			"IssueCredentialV3SendOffer",
			strings.ToUpper("Post"),
			"/agent/command/issue-credential-v3/send-offer",
			c.IssueCredentialV3SendOffer,
		},
		{
			"IssueCredentialV3SendProposal",
			strings.ToUpper("Post"),
			"/agent/command/issue-credential-v3/send-proposal",
			c.IssueCredentialV3SendProposal,
		},
		{
			"IssueCredentialV3SendRequest",
			strings.ToUpper("Post"),
			"/agent/command/issue-credential-v3/send-request",
			c.IssueCredentialV3SendRequest,
		},
		{
			"IssueCredentialV3Store",
			strings.ToUpper("Post"),
			"/agent/command/issue-credential-v3/store",
			c.IssueCredentialV3Store,
		},
	}
}

// IssueCredentialV3Issue - Issue credential
func (c *IssueCredentialV3ApiController) IssueCredentialV3Issue(w http.ResponseWriter, r *http.Request) {
	issueCredentialV3IssueRequestParam := IssueCredentialV3IssueRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&issueCredentialV3IssueRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIssueCredentialV3IssueRequestRequired(issueCredentialV3IssueRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.IssueCredentialV3Issue(r.Context(), issueCredentialV3IssueRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// IssueCredentialV3RetrieveCredentialApplication - Retrieves the Credential Application object received by the agent.
func (c *IssueCredentialV3ApiController) IssueCredentialV3RetrieveCredentialApplication(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.IssueCredentialV3RetrieveCredentialApplication(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// IssueCredentialV3RetrieveCredentialProposal - Retrieves the Propose Credential object received by the agent.
func (c *IssueCredentialV3ApiController) IssueCredentialV3RetrieveCredentialProposal(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.IssueCredentialV3RetrieveCredentialProposal(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// IssueCredentialV3SendOffer - Send credential offer
func (c *IssueCredentialV3ApiController) IssueCredentialV3SendOffer(w http.ResponseWriter, r *http.Request) {
	bodyParam := map[string]interface{}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.IssueCredentialV3SendOffer(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// IssueCredentialV3SendProposal - Send credential proposal
func (c *IssueCredentialV3ApiController) IssueCredentialV3SendProposal(w http.ResponseWriter, r *http.Request) {
	issueCredentialV3SendProposalRequestParam := IssueCredentialV3SendProposalRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&issueCredentialV3SendProposalRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIssueCredentialV3SendProposalRequestRequired(issueCredentialV3SendProposalRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.IssueCredentialV3SendProposal(r.Context(), issueCredentialV3SendProposalRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// IssueCredentialV3SendRequest - Send credential request
func (c *IssueCredentialV3ApiController) IssueCredentialV3SendRequest(w http.ResponseWriter, r *http.Request) {
	bodyParam := map[string]interface{}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.IssueCredentialV3SendRequest(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// IssueCredentialV3Store - Store/accept credential
func (c *IssueCredentialV3ApiController) IssueCredentialV3Store(w http.ResponseWriter, r *http.Request) {
	issueCredentialStoreRequestParam := IssueCredentialStoreRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&issueCredentialStoreRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIssueCredentialStoreRequestRequired(issueCredentialStoreRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.IssueCredentialV3Store(r.Context(), issueCredentialStoreRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}