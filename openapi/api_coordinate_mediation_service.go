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

// CoordinateMediationApiService is a service that implements the logic for the CoordinateMediationApiServicer
// This service should implement the business logic for every endpoint for the CoordinateMediationApi API.
// Include any external packages or services that will be required by this service.
type CoordinateMediationApiService struct {
}

// NewCoordinateMediationApiService creates a default api service
func NewCoordinateMediationApiService() CoordinateMediationApiServicer {
	return &CoordinateMediationApiService{}
}

// CoordinateMediationGetByConnectionId - Get mediation record by connection id
func (s *CoordinateMediationApiService) CoordinateMediationGetByConnectionId(ctx context.Context, connectionId string) (ImplResponse, error) {
	// TODO - update CoordinateMediationGetByConnectionId with the required logic for this service method.
	// Add api_coordinate_mediation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CoordinateMediationOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, CoordinateMediationOperationResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CoordinateMediationGetByConnectionId method not implemented")
}

// CoordinateMediationSendDeny - Send mediation deny message
func (s *CoordinateMediationApiService) CoordinateMediationSendDeny(ctx context.Context, connectionAcceptInvitationRequest ConnectionAcceptInvitationRequest) (ImplResponse, error) {
	// TODO - update CoordinateMediationSendDeny with the required logic for this service method.
	// Add api_coordinate_mediation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CoordinateMediationOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, CoordinateMediationOperationResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CoordinateMediationSendDeny method not implemented")
}

// CoordinateMediationSendGrant - Send mediation grant message
func (s *CoordinateMediationApiService) CoordinateMediationSendGrant(ctx context.Context, connectionAcceptInvitationRequest ConnectionAcceptInvitationRequest) (ImplResponse, error) {
	// TODO - update CoordinateMediationSendGrant with the required logic for this service method.
	// Add api_coordinate_mediation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CoordinateMediationOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, CoordinateMediationOperationResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CoordinateMediationSendGrant method not implemented")
}

// CoordinateMediationSendRequest - Send mediation request message
func (s *CoordinateMediationApiService) CoordinateMediationSendRequest(ctx context.Context, connectionAcceptInvitationRequest ConnectionAcceptInvitationRequest) (ImplResponse, error) {
	// TODO - update CoordinateMediationSendRequest with the required logic for this service method.
	// Add api_coordinate_mediation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CoordinateMediationOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, CoordinateMediationOperationResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CoordinateMediationSendRequest method not implemented")
}
