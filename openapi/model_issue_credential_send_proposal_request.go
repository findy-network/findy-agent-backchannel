/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IssueCredentialSendProposalRequest struct {

	Data IssueCredentialSendProposalRequestData `json:"data"`
}

// AssertIssueCredentialSendProposalRequestRequired checks if the required fields are not zero-ed
func AssertIssueCredentialSendProposalRequestRequired(obj IssueCredentialSendProposalRequest) error {
	elements := map[string]interface{}{
		"data": obj.Data,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertIssueCredentialSendProposalRequestDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertRecurseIssueCredentialSendProposalRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IssueCredentialSendProposalRequest (e.g. [][]IssueCredentialSendProposalRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIssueCredentialSendProposalRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIssueCredentialSendProposalRequest, ok := obj.(IssueCredentialSendProposalRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIssueCredentialSendProposalRequestRequired(aIssueCredentialSendProposalRequest)
	})
}