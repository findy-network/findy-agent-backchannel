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
	"encoding/json"
	"errors"
	"net/http"

	"github.com/findy-network/findy-agent-backchannel/agent"
)

// ConnectionApiService is a service that implements the logic for the ConnectionApiServicer
// This service should implement the business logic for every endpoint for the ConnectionApi API.
// Include any external packages or services that will be required by this service.
type ConnectionApiService struct {
	a *agent.Agent
}

// NewConnectionApiService creates a default api service
func NewConnectionApiService(a *agent.Agent) ConnectionApiServicer {
	return &ConnectionApiService{a: a}
}

// ConnectionAcceptInvitation - Accept an invitation
func (s *ConnectionApiService) ConnectionAcceptInvitation(ctx context.Context, connectionAcceptInvitationRequest ConnectionAcceptInvitationRequest) (ImplResponse, error) {
	id, err := s.a.RequestConnection(connectionAcceptInvitationRequest.Id)
	if err == nil {
		return Response(200, ConnectionAcceptInvitation200Response{ConnectionId: id, State: REQUEST}), nil
	}
	return Response(http.StatusInternalServerError, nil), err
}

// ConnectionAcceptRequest - Accept a connection request
func (s *ConnectionApiService) ConnectionAcceptRequest(ctx context.Context, connectionAcceptInvitationRequest ConnectionAcceptInvitationRequest) (ImplResponse, error) {
	// Findy Agency does not have accept connection step at the moment
	return Response(200, ConnectionAcceptRequest200Response{connectionAcceptInvitationRequest.Id, RESPONSE}), nil
}

// ConnectionCreateInvitation - Create a new connection invitation
func (s *ConnectionApiService) ConnectionCreateInvitation(ctx context.Context, connectionCreateInvitationRequest ConnectionCreateInvitationRequest) (ImplResponse, error) {
	invitationJSON, err := s.a.CreateInvitation()
	if err == nil {
		var invitationMap map[string]interface{}
		if err = json.Unmarshal([]byte(invitationJSON), &invitationMap); err == nil {
			return Response(200, ConnectionCreateInvitation200Response{ConnectionId: invitationMap["@id"].(string), Invitation: invitationMap}), nil
		}
	}
	return Response(http.StatusInternalServerError, nil), err
}

// ConnectionGetAll - Get all connections
func (s *ConnectionApiService) ConnectionGetAll(ctx context.Context) (ImplResponse, error) {
	// TODO - update ConnectionGetAll with the required logic for this service method.
	// Add api_connection_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []ConnectionResponse{}) or use other options such as http.Ok ...
	//return Response(200, []ConnectionResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConnectionGetAll method not implemented")
}

// ConnectionGetById - Get connection by id
func (s *ConnectionApiService) ConnectionGetById(ctx context.Context, connectionId string) (ImplResponse, error) {
	if _, err := s.a.GetConnection(connectionId); err == nil {
		return Response(200, ConnectionResponse{ConnectionId: connectionId, State: ACTIVE}), nil
	}
	if _, err := s.a.GetConnectionInvitation(connectionId); err == nil {
		return Response(200, ConnectionResponse{ConnectionId: connectionId, State: INVITATION}), nil
	}
	return Response(http.StatusNotFound, nil), nil
}

// ConnectionReceiveInvitation - Receive an invitation
func (s *ConnectionApiService) ConnectionReceiveInvitation(ctx context.Context, connectionReceiveInvitationRequest ConnectionReceiveInvitationRequest) (ImplResponse, error) {
	invitationBytes, err := json.Marshal(connectionReceiveInvitationRequest.Data)
	if err != nil {
		return Response(http.StatusBadRequest, nil), err
	}
	id := "TODO" //connectionReceiveInvitationRequest.Data["@id"].(string)
	id, err = s.a.AddConnectionInvitation(id, string(invitationBytes))
	if err == nil {
		return Response(200, ConnectionReceiveInvitation200Response{ConnectionId: id, State: INVITATION}), nil
	}
	return Response(http.StatusInternalServerError, nil), err
}

// ConnectionSendPing - Send trust ping
func (s *ConnectionApiService) ConnectionSendPing(ctx context.Context, connectionSendPingRequest ConnectionSendPingRequest) (ImplResponse, error) {
	id, err := s.a.TrustPing(connectionSendPingRequest.Id)
	if err == nil {
		return Response(200, ConnectionAcceptRequest200Response{ConnectionId: id, State: ACTIVE}), nil
	}
	return Response(http.StatusNotFound, nil), err
}
