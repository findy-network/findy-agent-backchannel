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

	"github.com/gorilla/mux"
)

// A PresentProofApiController binds http requests to an api service and writes the service results to the http response
type PresentProofApiController struct {
	service PresentProofApiServicer
}

// NewPresentProofApiController creates a default api controller
func NewPresentProofApiController(s PresentProofApiServicer) Router {
	return &PresentProofApiController{service: s}
}

// Routes returns all of the api route for the PresentProofApiController
func (c *PresentProofApiController) Routes() Routes {
	return Routes{ 
		{
			"PresentProofGetByThreadId",
			strings.ToUpper("Get"),
			"/agent/command/proof/{presentationExchangeThreadId}",
			c.PresentProofGetByThreadId,
		},
		{
			"PresentProofSendPresentation",
			strings.ToUpper("Post"),
			"/agent/command/proof/send-presentation",
			c.PresentProofSendPresentation,
		},
		{
			"PresentProofSendProposal",
			strings.ToUpper("Post"),
			"/agent/command/proof/send-proposal",
			c.PresentProofSendProposal,
		},
		{
			"PresentProofSendRequest",
			strings.ToUpper("Post"),
			"/agent/command/proof/send-request",
			c.PresentProofSendRequest,
		},
		{
			"PresentProofVerifyPresentation",
			strings.ToUpper("Post"),
			"/agent/command/proof/verify-presentation",
			c.PresentProofVerifyPresentation,
		},
	}
}

// PresentProofGetByThreadId - Get presentation exchange record by thread id
func (c *PresentProofApiController) PresentProofGetByThreadId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	presentationExchangeThreadId := params["presentationExchangeThreadId"]
	
	result, err := c.service.PresentProofGetByThreadId(r.Context(), presentationExchangeThreadId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PresentProofSendPresentation - Send presentation
func (c *PresentProofApiController) PresentProofSendPresentation(w http.ResponseWriter, r *http.Request) {
	inlineObject9 := &InlineObject9{}
	if err := json.NewDecoder(r.Body).Decode(&inlineObject9); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.PresentProofSendPresentation(r.Context(), *inlineObject9)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PresentProofSendProposal - Send presentation proposal
func (c *PresentProofApiController) PresentProofSendProposal(w http.ResponseWriter, r *http.Request) {
	inlineObject7 := &InlineObject7{}
	if err := json.NewDecoder(r.Body).Decode(&inlineObject7); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.PresentProofSendProposal(r.Context(), *inlineObject7)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PresentProofSendRequest - Send presentation request
func (c *PresentProofApiController) PresentProofSendRequest(w http.ResponseWriter, r *http.Request) {
	inlineObject8 := &InlineObject8{}
	if err := json.NewDecoder(r.Body).Decode(&inlineObject8); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.PresentProofSendRequest(r.Context(), *inlineObject8)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PresentProofVerifyPresentation - Verify presentation
func (c *PresentProofApiController) PresentProofVerifyPresentation(w http.ResponseWriter, r *http.Request) {
	inlineObject10 := &InlineObject10{}
	if err := json.NewDecoder(r.Body).Decode(&inlineObject10); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.PresentProofVerifyPresentation(r.Context(), *inlineObject10)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
