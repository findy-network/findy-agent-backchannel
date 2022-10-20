/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PresentProofSendProposalRequest struct {

	Data PresentProofSendProposalRequestData `json:"data"`
}

// AssertPresentProofSendProposalRequestRequired checks if the required fields are not zero-ed
func AssertPresentProofSendProposalRequestRequired(obj PresentProofSendProposalRequest) error {
	elements := map[string]interface{}{
		"data": obj.Data,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPresentProofSendProposalRequestDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertRecursePresentProofSendProposalRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PresentProofSendProposalRequest (e.g. [][]PresentProofSendProposalRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePresentProofSendProposalRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPresentProofSendProposalRequest, ok := obj.(PresentProofSendProposalRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPresentProofSendProposalRequestRequired(aPresentProofSendProposalRequest)
	})
}
