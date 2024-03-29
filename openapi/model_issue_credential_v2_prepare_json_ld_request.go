/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IssueCredentialV2PrepareJsonLdRequest struct {

	DidMethod string `json:"did_method"`

	ProofType string `json:"proof_type"`
}

// AssertIssueCredentialV2PrepareJsonLdRequestRequired checks if the required fields are not zero-ed
func AssertIssueCredentialV2PrepareJsonLdRequestRequired(obj IssueCredentialV2PrepareJsonLdRequest) error {
	elements := map[string]interface{}{
		"did_method": obj.DidMethod,
		"proof_type": obj.ProofType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseIssueCredentialV2PrepareJsonLdRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IssueCredentialV2PrepareJsonLdRequest (e.g. [][]IssueCredentialV2PrepareJsonLdRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIssueCredentialV2PrepareJsonLdRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIssueCredentialV2PrepareJsonLdRequest, ok := obj.(IssueCredentialV2PrepareJsonLdRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIssueCredentialV2PrepareJsonLdRequestRequired(aIssueCredentialV2PrepareJsonLdRequest)
	})
}
