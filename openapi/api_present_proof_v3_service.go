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

// PresentProofV3ApiService is a service that implements the logic for the PresentProofV3ApiServicer
// This service should implement the business logic for every endpoint for the PresentProofV3Api API.
// Include any external packages or services that will be required by this service.
type PresentProofV3ApiService struct {
}

// NewPresentProofV3ApiService creates a default api service
func NewPresentProofV3ApiService() PresentProofV3ApiServicer {
	return &PresentProofV3ApiService{}
}

// PresentProofV3SendPresentation - Send presentation
func (s *PresentProofV3ApiService) PresentProofV3SendPresentation(ctx context.Context, presentProofV2SendPresentationRequest PresentProofV2SendPresentationRequest) (ImplResponse, error) {
	// TODO - update PresentProofV3SendPresentation with the required logic for this service method.
	// Add api_present_proof_v3_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofSendPresentation200Response{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofSendPresentation200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofV3SendPresentation method not implemented")
}

// PresentProofV3SendProposal - Send presentation proposal
func (s *PresentProofV3ApiService) PresentProofV3SendProposal(ctx context.Context, presentProofV3SendProposalRequest PresentProofV3SendProposalRequest) (ImplResponse, error) {
	// TODO - update PresentProofV3SendProposal with the required logic for this service method.
	// Add api_present_proof_v3_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofSendProposal200Response{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofSendProposal200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofV3SendProposal method not implemented")
}

// PresentProofV3SendRequest - Send present-proof v3 presentation request
func (s *PresentProofV3ApiService) PresentProofV3SendRequest(ctx context.Context, presentProofV3SendRequestRequest PresentProofV3SendRequestRequest) (ImplResponse, error) {
	// TODO - update PresentProofV3SendRequest with the required logic for this service method.
	// Add api_present_proof_v3_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofSendRequest200Response{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofSendRequest200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofV3SendRequest method not implemented")
}

// PresentProofV3VerifyPresentation - Verify presentation
func (s *PresentProofV3ApiService) PresentProofV3VerifyPresentation(ctx context.Context, presentProofVerifyPresentationRequest PresentProofVerifyPresentationRequest) (ImplResponse, error) {
	// TODO - update PresentProofV3VerifyPresentation with the required logic for this service method.
	// Add api_present_proof_v3_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofV2VerifyPresentation200Response{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofV2VerifyPresentation200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofV3VerifyPresentation method not implemented")
}
