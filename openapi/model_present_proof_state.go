/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PresentProofState : All possible present proof state values
type PresentProofState string

// List of PresentProofState
const (
	PROOF_PROPOSAL_SENT PresentProofState = "proposal-sent"
	PROOF_PROPOSAL_RECEIVED PresentProofState = "proposal-received"
	PROOF_REQUEST_SENT PresentProofState = "request-sent"
	PROOF_REQUEST_RECEIVED PresentProofState = "request-received"
	PROOF_PRESENTATION_SENT PresentProofState = "presentation-sent"
	PROOF_PRESENTATION_RECEIVED PresentProofState = "presentation-received"
	PROOF_REJECT_SENT PresentProofState = "reject-sent"
	PROOF_DONE PresentProofState = "done"
)
