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
	"log"
	"net/http"

	"github.com/findy-network/findy-agent-backchannel/agent"
)

// CredentialApiService is a service that implents the logic for the CredentialApiServicer
// This service should implement the business logic for every endpoint for the CredentialApi API.
// Include any external packages or services that will be required by this service.
type CredentialApiService struct {
	a *agent.Agent
}

// NewCredentialApiService creates a default api service
func NewCredentialApiService(a *agent.Agent) CredentialApiServicer {
	return &CredentialApiService{
		a: a,
	}
}

// CredentialGetById - Get credential by id
func (s *CredentialApiService) CredentialGetById(ctx context.Context, credentialId string) (ImplResponse, error) {
	cred, err := s.a.GetCredentialContent(credentialId)
	log.Println("Credential content", cred, credentialId)
	if err != nil {
		return Response(http.StatusNotFound, nil), err
	}
	return Response(200, cred), nil
}
