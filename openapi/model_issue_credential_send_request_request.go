/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IssueCredentialSendRequestRequest struct {

	Id string `json:"id"`
}

// AssertIssueCredentialSendRequestRequestRequired checks if the required fields are not zero-ed
func AssertIssueCredentialSendRequestRequestRequired(obj IssueCredentialSendRequestRequest) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseIssueCredentialSendRequestRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IssueCredentialSendRequestRequest (e.g. [][]IssueCredentialSendRequestRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIssueCredentialSendRequestRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIssueCredentialSendRequestRequest, ok := obj.(IssueCredentialSendRequestRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIssueCredentialSendRequestRequestRequired(aIssueCredentialSendRequestRequest)
	})
}
