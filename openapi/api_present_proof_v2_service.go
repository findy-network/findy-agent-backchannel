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

// PresentProofV2ApiService is a service that implements the logic for the PresentProofV2ApiServicer
// This service should implement the business logic for every endpoint for the PresentProofV2Api API.
// Include any external packages or services that will be required by this service.
type PresentProofV2ApiService struct {
}

// NewPresentProofV2ApiService creates a default api service
func NewPresentProofV2ApiService() PresentProofV2ApiServicer {
	return &PresentProofV2ApiService{}
}

// PresentProofV2SendPresentation - Send presentation
func (s *PresentProofV2ApiService) PresentProofV2SendPresentation(ctx context.Context, presentProofV2SendPresentationRequest PresentProofV2SendPresentationRequest) (ImplResponse, error) {
	// TODO - update PresentProofV2SendPresentation with the required logic for this service method.
	// Add api_present_proof_v2_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofSendPresentation200Response{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofSendPresentation200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofV2SendPresentation method not implemented")
}

// PresentProofV2SendRequest - Send presentation request
func (s *PresentProofV2ApiService) PresentProofV2SendRequest(ctx context.Context, presentProofV2SendRequestRequest PresentProofV2SendRequestRequest) (ImplResponse, error) {
	// TODO - update PresentProofV2SendRequest with the required logic for this service method.
	// Add api_present_proof_v2_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofSendRequest200Response{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofSendRequest200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofV2SendRequest method not implemented")
}

// PresentProofV2VerifyPresentation - Verify presentation
func (s *PresentProofV2ApiService) PresentProofV2VerifyPresentation(ctx context.Context, presentProofVerifyPresentationRequest PresentProofVerifyPresentationRequest) (ImplResponse, error) {
	// TODO - update PresentProofV2VerifyPresentation with the required logic for this service method.
	// Add api_present_proof_v2_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofV2VerifyPresentation200Response{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofV2VerifyPresentation200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofV2VerifyPresentation method not implemented")
}
