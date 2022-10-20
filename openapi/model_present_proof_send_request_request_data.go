/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofPresentationRequestProofRequest struct {
	Data PresentProofSendPresentationRequestData `json:"data"`
}

type PresentProofPresentationRequest struct {
	Comment      string                                      `json:"comment"`
	ProofRequest PresentProofPresentationRequestProofRequest `json:"proof_request"`
}

type PresentProofSendRequestRequestData struct {
	ConnectionId string `json:"connection_id"`

	PresentationRequest PresentProofPresentationRequest `json:"presentation_request"`
}

// AssertPresentProofSendRequestRequestDataRequired checks if the required fields are not zero-ed
func AssertPresentProofSendRequestRequestDataRequired(obj PresentProofSendRequestRequestData) error {
	elements := map[string]interface{}{
		"connection_id":        obj.ConnectionId,
		"presentation_request": obj.PresentationRequest,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecursePresentProofSendRequestRequestDataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofSendRequestRequestData (e.g. [][]PresentProofSendRequestRequestData), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofSendRequestRequestDataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofSendRequestRequestData, ok := obj.(PresentProofSendRequestRequestData)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofSendRequestRequestDataRequired(aPresentProofSendRequestRequestData)
	})
}