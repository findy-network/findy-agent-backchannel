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
)

var (
	Version = "dev"
)

// StatusApiService is a service that implements the logic for the StatusApiServicer
// This service should implement the business logic for every endpoint for the StatusApi API.
// Include any external packages or services that will be required by this service.
type StatusApiService struct {
}

// NewStatusApiService creates a default api service
func NewStatusApiService() StatusApiServicer {
	return &StatusApiService{}
}

// StatusGet - Get agent/backchannel status
func (s *StatusApiService) StatusGet(ctx context.Context) (ImplResponse, error) {
	return Response(200, StatusGet200Response{Status: string(ACTIVE)}), nil
}

// VersionGet - Get agent/backchannel version
func (s *StatusApiService) VersionGet(ctx context.Context) (ImplResponse, error) {
	return Response(200, Version), nil
}
