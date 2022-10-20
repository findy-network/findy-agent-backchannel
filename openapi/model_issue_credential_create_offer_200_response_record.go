/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IssueCredentialCreateOffer200ResponseRecord struct {

	State *interface{} `json:"state"`

	ThreadId string `json:"thread_id"`

	CredentialId string `json:"credential_id,omitempty"`
}

// AssertIssueCredentialCreateOffer200ResponseRecordRequired checks if the required fields are not zero-ed
func AssertIssueCredentialCreateOffer200ResponseRecordRequired(obj IssueCredentialCreateOffer200ResponseRecord) error {
	elements := map[string]interface{}{
		"state": obj.State,
		"thread_id": obj.ThreadId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseIssueCredentialCreateOffer200ResponseRecordRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IssueCredentialCreateOffer200ResponseRecord (e.g. [][]IssueCredentialCreateOffer200ResponseRecord), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIssueCredentialCreateOffer200ResponseRecordRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIssueCredentialCreateOffer200ResponseRecord, ok := obj.(IssueCredentialCreateOffer200ResponseRecord)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIssueCredentialCreateOffer200ResponseRecordRequired(aIssueCredentialCreateOffer200ResponseRecord)
	})
}