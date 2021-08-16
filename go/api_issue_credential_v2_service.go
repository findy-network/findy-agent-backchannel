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
	"context"
	"net/http"
	"errors"
)

// IssueCredentialV2ApiService is a service that implents the logic for the IssueCredentialV2ApiServicer
// This service should implement the business logic for every endpoint for the IssueCredentialV2Api API. 
// Include any external packages or services that will be required by this service.
type IssueCredentialV2ApiService struct {
}

// NewIssueCredentialV2ApiService creates a default api service
func NewIssueCredentialV2ApiService() IssueCredentialV2ApiServicer {
	return &IssueCredentialV2ApiService{}
}

// IssueCredentialV2PrepareJsonLD - Prepare for issuing a JSON-LD credential (RFC0593)
func (s *IssueCredentialV2ApiService) IssueCredentialV2PrepareJsonLD(ctx context.Context, uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) (ImplResponse, error) {
	// TODO - update IssueCredentialV2PrepareJsonLD with the required logic for this service method.
	// Add api_issue_credential_v2_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("IssueCredentialV2PrepareJsonLD method not implemented")
}

