/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/findy-network/findy-agent-backchannel/openapi"
)

func main() {
	log.Printf("Server started")

	AgentApiService := openapi.NewAgentApiService()
	AgentApiController := openapi.NewAgentApiController(AgentApiService)

	ConnectionApiService := openapi.NewConnectionApiService()
	ConnectionApiController := openapi.NewConnectionApiController(ConnectionApiService)

	CoordinateMediationApiService := openapi.NewCoordinateMediationApiService()
	CoordinateMediationApiController := openapi.NewCoordinateMediationApiController(CoordinateMediationApiService)

	CredentialApiService := openapi.NewCredentialApiService()
	CredentialApiController := openapi.NewCredentialApiController(CredentialApiService)

	CredentialDefinitionApiService := openapi.NewCredentialDefinitionApiService()
	CredentialDefinitionApiController := openapi.NewCredentialDefinitionApiController(CredentialDefinitionApiService)

	DIDApiService := openapi.NewDIDApiService()
	DIDApiController := openapi.NewDIDApiController(DIDApiService)

	DIDExchangeApiService := openapi.NewDIDExchangeApiService()
	DIDExchangeApiController := openapi.NewDIDExchangeApiController(DIDExchangeApiService)

	IssueCredentialApiService := openapi.NewIssueCredentialApiService()
	IssueCredentialApiController := openapi.NewIssueCredentialApiController(IssueCredentialApiService)

	IssueCredentialV2ApiService := openapi.NewIssueCredentialV2ApiService()
	IssueCredentialV2ApiController := openapi.NewIssueCredentialV2ApiController(IssueCredentialV2ApiService)

	IssueCredentialV3ApiService := openapi.NewIssueCredentialV3ApiService()
	IssueCredentialV3ApiController := openapi.NewIssueCredentialV3ApiController(IssueCredentialV3ApiService)

	OutOfBandApiService := openapi.NewOutOfBandApiService()
	OutOfBandApiController := openapi.NewOutOfBandApiController(OutOfBandApiService)

	OutOfBandV2ApiService := openapi.NewOutOfBandV2ApiService()
	OutOfBandV2ApiController := openapi.NewOutOfBandV2ApiController(OutOfBandV2ApiService)

	PresentProofApiService := openapi.NewPresentProofApiService()
	PresentProofApiController := openapi.NewPresentProofApiController(PresentProofApiService)

	PresentProofV2ApiService := openapi.NewPresentProofV2ApiService()
	PresentProofV2ApiController := openapi.NewPresentProofV2ApiController(PresentProofV2ApiService)

	PresentProofV3ApiService := openapi.NewPresentProofV3ApiService()
	PresentProofV3ApiController := openapi.NewPresentProofV3ApiController(PresentProofV3ApiService)

	RevocationApiService := openapi.NewRevocationApiService()
	RevocationApiController := openapi.NewRevocationApiController(RevocationApiService)

	SchemaApiService := openapi.NewSchemaApiService()
	SchemaApiController := openapi.NewSchemaApiController(SchemaApiService)

	StatusApiService := openapi.NewStatusApiService()
	StatusApiController := openapi.NewStatusApiController(StatusApiService)

	router := openapi.NewRouter(AgentApiController, ConnectionApiController, CoordinateMediationApiController, CredentialApiController, CredentialDefinitionApiController, DIDApiController, DIDExchangeApiController, IssueCredentialApiController, IssueCredentialV2ApiController, IssueCredentialV3ApiController, OutOfBandApiController, OutOfBandV2ApiController, PresentProofApiController, PresentProofV2ApiController, PresentProofV3ApiController, RevocationApiController, SchemaApiController, StatusApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
