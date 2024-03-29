/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IssueCredentialV2CredentialProposalFilter struct {

	Indy IssueCredentialV2AttachFormatIndyCredFilter `json:"indy,omitempty"`

	JsonLd map[string]interface{} `json:"json-ld,omitempty"`
}

// AssertIssueCredentialV2CredentialProposalFilterRequired checks if the required fields are not zero-ed
func AssertIssueCredentialV2CredentialProposalFilterRequired(obj IssueCredentialV2CredentialProposalFilter) error {
	if err := AssertIssueCredentialV2AttachFormatIndyCredFilterRequired(obj.Indy); err != nil {
		return err
	}
	return nil
}

// AssertRecurseIssueCredentialV2CredentialProposalFilterRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IssueCredentialV2CredentialProposalFilter (e.g. [][]IssueCredentialV2CredentialProposalFilter), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIssueCredentialV2CredentialProposalFilterRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIssueCredentialV2CredentialProposalFilter, ok := obj.(IssueCredentialV2CredentialProposalFilter)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIssueCredentialV2CredentialProposalFilterRequired(aIssueCredentialV2CredentialProposalFilter)
	})
}
