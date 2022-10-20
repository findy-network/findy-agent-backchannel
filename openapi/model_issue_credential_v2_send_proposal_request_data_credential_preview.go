/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IssueCredentialV2SendProposalRequestDataCredentialPreview struct {

	Type *interface{} `json:"@type"`

	Attributes []CredentialPreviewAttributesInner `json:"attributes"`
}

// AssertIssueCredentialV2SendProposalRequestDataCredentialPreviewRequired checks if the required fields are not zero-ed
func AssertIssueCredentialV2SendProposalRequestDataCredentialPreviewRequired(obj IssueCredentialV2SendProposalRequestDataCredentialPreview) error {
	elements := map[string]interface{}{
		"@type": obj.Type,
		"attributes": obj.Attributes,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Attributes {
		if err := AssertCredentialPreviewAttributesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseIssueCredentialV2SendProposalRequestDataCredentialPreviewRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IssueCredentialV2SendProposalRequestDataCredentialPreview (e.g. [][]IssueCredentialV2SendProposalRequestDataCredentialPreview), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIssueCredentialV2SendProposalRequestDataCredentialPreviewRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIssueCredentialV2SendProposalRequestDataCredentialPreview, ok := obj.(IssueCredentialV2SendProposalRequestDataCredentialPreview)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIssueCredentialV2SendProposalRequestDataCredentialPreviewRequired(aIssueCredentialV2SendProposalRequestDataCredentialPreview)
	})
}