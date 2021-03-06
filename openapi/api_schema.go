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

// A SchemaApiController binds http requests to an api service and writes the service results to the http response
type SchemaApiController struct {
	service SchemaApiServicer
}

// NewSchemaApiController creates a default api controller
func NewSchemaApiController(s SchemaApiServicer) Router {
	return &SchemaApiController{service: s}
}

// Routes returns all of the api route for the SchemaApiController
func (c *SchemaApiController) Routes() Routes {
	return Routes{
		{
			"SchemaCreate",
			strings.ToUpper("Post"),
			"/agent/command/{schema:schema\\/?}",
			c.SchemaCreate,
		},
		{
			"SchemaGetById",
			strings.ToUpper("Get"),
			"/agent/command/schema/{schemaId}",
			c.SchemaGetById,
		},
	}
}

// SchemaCreate - Create a new schema
func (c *SchemaApiController) SchemaCreate(w http.ResponseWriter, r *http.Request) {
	inlineObject4 := &InlineObject4{}
	if err := json.NewDecoder(r.Body).Decode(&inlineObject4); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.SchemaCreate(r.Context(), *inlineObject4)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// SchemaGetById - Get schema by id
func (c *SchemaApiController) SchemaGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	schemaId := params["schemaId"]

	result, err := c.service.SchemaGetById(r.Context(), schemaId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
