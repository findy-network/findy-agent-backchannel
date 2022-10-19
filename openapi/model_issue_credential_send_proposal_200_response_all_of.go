/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IssueCredentialSendProposal200ResponseAllOf struct {

	State *interface{} `json:"state,omitempty"`
}

// AssertIssueCredentialSendProposal200ResponseAllOfRequired checks if the required fields are not zero-ed
func AssertIssueCredentialSendProposal200ResponseAllOfRequired(obj IssueCredentialSendProposal200ResponseAllOf) error {
	return nil
}

// AssertRecurseIssueCredentialSendProposal200ResponseAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IssueCredentialSendProposal200ResponseAllOf (e.g. [][]IssueCredentialSendProposal200ResponseAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIssueCredentialSendProposal200ResponseAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIssueCredentialSendProposal200ResponseAllOf, ok := obj.(IssueCredentialSendProposal200ResponseAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIssueCredentialSendProposal200ResponseAllOfRequired(aIssueCredentialSendProposal200ResponseAllOf)
	})
}
