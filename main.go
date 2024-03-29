/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.
 * For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/findy-network/findy-agent-backchannel/agent"
	openapi "github.com/findy-network/findy-agent-backchannel/openapi"
	"github.com/gorilla/mux"
	"github.com/lainio/err2/try"
)

func createRouter(a *agent.Agent) *mux.Router {
	AgentApiService := openapi.NewAgentApiService()
	AgentApiController := openapi.NewAgentApiController(AgentApiService)

	ConnectionApiService := openapi.NewConnectionApiService(a)
	ConnectionApiController := openapi.NewConnectionApiController(ConnectionApiService)

	CoordinateMediationApiService := openapi.NewCoordinateMediationApiService()
	CoordinateMediationApiController := openapi.NewCoordinateMediationApiController(CoordinateMediationApiService)

	CredentialApiService := openapi.NewCredentialApiService(a)
	CredentialApiController := openapi.NewCredentialApiController(CredentialApiService)

	CredentialDefinitionApiService := openapi.NewCredentialDefinitionApiService(a)
	CredentialDefinitionApiController := openapi.NewCredentialDefinitionApiController(CredentialDefinitionApiService)

	DIDApiService := openapi.NewDIDApiService(a)
	DIDApiController := openapi.NewDIDApiController(DIDApiService)

	DIDExchangeApiService := openapi.NewDIDExchangeApiService(a)
	DIDExchangeApiController := openapi.NewDIDExchangeApiController(DIDExchangeApiService)

	IssueCredentialApiService := openapi.NewIssueCredentialApiService(a)
	IssueCredentialApiController := openapi.NewIssueCredentialApiController(IssueCredentialApiService)

	IssueCredentialV2ApiService := openapi.NewIssueCredentialV2ApiService()
	IssueCredentialV2ApiController := openapi.NewIssueCredentialV2ApiController(IssueCredentialV2ApiService)

	IssueCredentialV3ApiService := openapi.NewIssueCredentialV3ApiService()
	IssueCredentialV3ApiController := openapi.NewIssueCredentialV3ApiController(IssueCredentialV3ApiService)

	OutOfBandApiService := openapi.NewOutOfBandApiService(a)
	OutOfBandApiController := openapi.NewOutOfBandApiController(OutOfBandApiService)

	OutOfBandV2ApiService := openapi.NewOutOfBandV2ApiService()
	OutOfBandV2ApiController := openapi.NewOutOfBandV2ApiController(OutOfBandV2ApiService)

	PresentProofApiService := openapi.NewPresentProofApiService(a)
	PresentProofApiController := openapi.NewPresentProofApiController(PresentProofApiService)

	PresentProofV2ApiService := openapi.NewPresentProofV2ApiService()
	PresentProofV2ApiController := openapi.NewPresentProofV2ApiController(PresentProofV2ApiService)

	PresentProofV3ApiService := openapi.NewPresentProofV3ApiService()
	PresentProofV3ApiController := openapi.NewPresentProofV3ApiController(PresentProofV3ApiService)

	RevocationApiService := openapi.NewRevocationApiService()
	RevocationApiController := openapi.NewRevocationApiController(RevocationApiService)

	SchemaApiService := openapi.NewSchemaApiService(a)
	SchemaApiController := openapi.NewSchemaApiController(SchemaApiService)

	StatusApiService := openapi.NewStatusApiService()
	StatusApiController := openapi.NewStatusApiController(StatusApiService)

	return openapi.NewRouter(
		AgentApiController,
		ConnectionApiController,
		CoordinateMediationApiController,
		CredentialApiController,
		CredentialDefinitionApiController,
		DIDApiController,
		DIDExchangeApiController,
		IssueCredentialApiController,
		IssueCredentialV2ApiController,
		IssueCredentialV3ApiController,
		OutOfBandApiController,
		OutOfBandV2ApiController,
		PresentProofApiController,
		PresentProofV2ApiController,
		PresentProofV3ApiController,
		RevocationApiController,
		SchemaApiController,
		StatusApiController,
	)
}

func main() {
	port := "9999"
	if strings.HasPrefix(os.Args[2], "9") {
		port = os.Args[2]
	}

	log.Printf("Starting server at port %s", port)

	a := agent.Init()
	a.Login()

	log.Printf("Agent login succeeded")

	router := createRouter(a)

	var handler http.Handler = router
	if os.Getenv("FAB_LOG_INCOMING_REQUESTS") == "true" {
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bodyBytes, _ := io.ReadAll(r.Body)
			try.To(r.Body.Close())
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			log.Println(r.Method + " " + r.URL.String() + " " + string(bodyBytes))
			router.ServeHTTP(w, r)
		})
	}

	server := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           handler,
	}

	try.To(server.ListenAndServe())
}
